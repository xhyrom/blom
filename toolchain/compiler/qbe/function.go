package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileFunction(declaration *ast.FunctionDeclaration) {
	c.Scopes.Append()

	arguments := make([]qbe.TypedValue, len(declaration.Arguments))
	for i, argument := range declaration.Arguments {
		t := qbe.RemapAstType(argument.Type)

		temp := c.createVariable(t, argument.Name)
		arguments[i] = qbe.NewTypedValue(t, temp)
	}

	var linkage qbe.Linkage
	if declaration.HasAnnotation(ast.Public) {
		linkage = qbe.NewLinkage(true)
	} else {
		linkage = qbe.NewLinkage(false)
	}

	returnType := qbe.RemapAstType(declaration.ReturnType)
	function := qbe.Function{
		Linkage:    linkage,
		Name:       declaration.Name,
		Arguments:  arguments,
		ReturnType: returnType,
		Variadic:   declaration.Variadic,
		External:   declaration.IsNative(),
		Blocks:     make([]qbe.Block, 0),
	}

	if declaration.IsNative() {
		c.Scopes.Pop()
		return
	}

	function.AddBlock("start")

	for _, statement := range declaration.Body {
		c.compileStatement(statement, &function, nil, false)
	}

	c.Scopes.Pop()

	c.Module.SetFunctionByName(declaration.Name, function)
}

func (c *Compiler) compileFunctionCall(call *ast.FunctionCall, currentFunction *qbe.Function, vtype *qbe.Type) *qbe.TypedValue {
	function := c.Module.GetFunctionByName(call.Name)

	parameters := make([]qbe.TypedValue, 0)
	for i, parameter := range call.Parameters {
		var argType *qbe.Type
		if i < len(function.Arguments) {
			argType = &function.Arguments[i].Type
		} else {
			argType = vtype
		}

		value := *c.compileStatement(parameter, currentFunction, argType, false)

		// Promotes f32 to f64 acording to the ISO C standard
		// https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
		if value.Type == qbe.Single && i >= len(function.Arguments) {
			value = *c.convertToType(value.Type, qbe.Double, value.Value, currentFunction)
		}

		if len(function.Arguments) == i && function.Variadic {
			parameters = append(parameters, qbe.TypedValue{
				Value: qbe.NewLiteralValue("..."),
				Type:  qbe.Null,
			})
		}

		parameters = append(parameters, value)
	}

	tempValue := c.getTemporaryValue(nil)

	currentFunction.LastBlock().AddAssign(
		tempValue,
		function.ReturnType,
		qbe.NewCallInstruction(qbe.NewGlobalValue(function.Name), parameters...),
	)

	return &qbe.TypedValue{
		Type:  function.ReturnType,
		Value: tempValue,
	}
}
