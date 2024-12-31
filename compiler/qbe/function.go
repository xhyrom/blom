package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/scope"
)

func (c *Compiler) compileFunction(declaration *ast.FunctionDeclaration) qbe.Function {
	c.Scopes = append(c.Scopes, scope.New[*qbe.TypedValue]())

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
		c.Scopes = c.Scopes[:len(c.Scopes)-1]
		return function
	}

	function.AddBlock("start")

	for _, statement := range declaration.Body {
		c.compileStatement(statement, &function, nil, false)
	}

	c.Scopes = c.Scopes[:len(c.Scopes)-1]

	return function
}

func (c *Compiler) compileFunctionCall(call *ast.FunctionCall, currentFunction *qbe.Function, vtype *qbe.Type) *qbe.TypedValue {
	var function *qbe.Function
	if call.Name == currentFunction.Name {
		function = currentFunction
	} else {
		function = c.Module.GetFunctionByName(call.Name)
	}

	parameters := make([]qbe.TypedValue, 0)
	for i, parameter := range call.Parameters {
		var argType *qbe.Type
		if i < len(function.Arguments) {
			argType = &function.Arguments[i].Type
		} else {
			argType = vtype
		}

		if len(function.Arguments) == i && function.Variadic {
			parameters = append(parameters, qbe.TypedValue{
				Value: qbe.NewLiteralValue("..."),
				Type:  qbe.Null,
			})
		}

		parameters = append(parameters, *c.compileStatement(parameter, currentFunction, argType, false))
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
