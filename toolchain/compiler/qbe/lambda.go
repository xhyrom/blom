package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileLambda(declaration *ast.LambdaDeclaration, function *qbe.Function, vtype *qbe.Type) *qbe.TypedValue {
	c.Scopes.Append()

	arguments := make([]qbe.TypedValue, len(declaration.Arguments))
	for i, argument := range declaration.Arguments {
		t := qbe.RemapAstType(argument.Type)

		temp := c.createVariable(t, argument.Name)
		arguments[i] = qbe.NewTypedValue(t, temp)
	}

	tempValue := c.createGlobalVariable(qbe.Func, "lambda")
	returnType := qbe.RemapAstType(declaration.ReturnType)

	lambda := qbe.Function{
		Linkage:    qbe.NewLinkage(false),
		Name:       tempValue.Name,
		Arguments:  arguments,
		ReturnType: returnType,
		Variadic:   declaration.Variadic,
		Blocks:     make([]qbe.Block, 0),
	}

	lambda.AddBlock("start")

	for _, statement := range declaration.Body {
		c.compileStatement(statement, &lambda, nil, false)
	}

	c.Scopes.Pop()

	c.Module.AddFunction(lambda)

	return &qbe.TypedValue{
		Type:  qbe.NewFunctionBox(lambda),
		Value: tempValue,
	}
}
