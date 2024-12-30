package qbe

import (
	"blom/ast"
	"blom/env"
	"blom/qbe"
)

func (c *Compiler) compileFunction(declaration *ast.FunctionDeclaration) qbe.Function {
	c.Scopes = append(c.Scopes, env.New[*qbe.TypedValue]())

	arguments := make([]qbe.TypedValue, len(declaration.Arguments))
	// TODO add arguments to scope

	returnType := qbe.RemapAstType(declaration.ReturnType)
	function := qbe.Function{
		Linkage:    qbe.NewLinkage(true), // TODO: Implement linkage
		Name:       declaration.Name,
		Arguments:  arguments,
		ReturnType: &returnType,
		Variadic:   declaration.Variadic,
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

func (c *Compiler) compileFunctionCall(call *ast.FunctionCall, currentFunction *qbe.Function) *qbe.TypedValue {
	var function *qbe.Function
	if call.Name == currentFunction.Name {
		function = currentFunction
	} else {
		function = c.Module.GetFunctionByName(call.Name)
	}

	tempValue := c.newTemporaryValue(nil)

	currentFunction.LastBlock().AddAssign(
		tempValue,
		*function.ReturnType,
		qbe.NewCallInstruction(qbe.NewGlobalValue(function.Name)),
	)

	return &qbe.TypedValue{
		Type:  *function.ReturnType,
		Value: tempValue,
	}
}
