package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileFunction(declaration *ast.FunctionDeclaration) {
	c.Scopes.Append()

	params := make([]qbe.TypedValue, len(declaration.Params))
	for i, param := range declaration.Params {
		t := qbe.RemapAstType(param.Type)

		temp := c.createVariable(t, param.Name.Value)
		params[i] = qbe.NewTypedValue(t, temp)
	}

	var linkage qbe.Linkage
	if declaration.HasAnnotation(ast.Public) {
		linkage = qbe.NewLinkage(true)
	} else {
		linkage = qbe.NewLinkage(false)
	}

	returnType := qbe.RemapAstType(declaration.Return)
	function := qbe.Function{
		Linkage:    linkage,
		Name:       declaration.Name.Value,
		Params:     params,
		ReturnType: returnType,
		Variadic:   declaration.HasAnnotation(ast.Variadic),
		External:   declaration.HasAnnotation(ast.Native),
		Blocks:     make([]qbe.Block, 0),
	}

	if declaration.HasAnnotation(ast.Native) {
		c.Scopes.Pop()
		return
	}

	function.AddBlock("start")

	for _, statement := range declaration.Block.Body {
		c.compileStatement(statement, &function, nil, false)
	}

	c.Scopes.Pop()

	c.Module.SetFunctionByName(declaration.Name.Value, function)
}

func (c *Compiler) compileFunctionCall(call *ast.FunctionCall, currentFunction *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	function := c.Module.GetFunctionByName(call.Path.Dotify())
	name := qbe.NewGlobalValue(function.Name)

	arguments := make([]qbe.TypedValue, 0)
	for i, arg := range call.Args {
		var paramType qbe.Type
		if i < len(function.Params) {
			paramType = function.Params[i].Type
		} else {
			paramType = vtype
		}

		value := *c.compileStatement(arg, currentFunction, paramType, false)

		// Promotes f32 to f64 acording to the ISO C standard
		// https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
		if value.Type == qbe.Single && i >= len(function.Params) {
			value = *c.convertToType(value.Type, qbe.Double, value.Value, currentFunction)
		}

		if len(function.Params) == i && function.Variadic {
			arguments = append(arguments, qbe.TypedValue{
				Value: qbe.NewLiteralValue("..."),
				Type:  qbe.Null,
			})
		}

		arguments = append(arguments, value)
	}

	tempValue := c.getTemporaryValue(nil)

	currentFunction.LastBlock().AddAssign(
		tempValue,
		function.ReturnType,
		qbe.NewCallInstruction(name, arguments...),
	)

	return &qbe.TypedValue{
		Type:  function.ReturnType,
		Value: tempValue,
	}
}
