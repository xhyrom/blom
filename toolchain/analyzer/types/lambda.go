package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeLambdaDeclaration(lambda *ast.LambdaDeclaration) ast.Type {
	a.Scopes.Append()

	for _, arg := range lambda.Arguments {
		a.Scopes.Set(arg.Name, &Variable{Type: arg.Type})
	}

	for _, statement := range lambda.Body {
		if statement.Kind() == ast.ReturnNode {
			ret := statement.(*ast.ReturnStatement)
			returnType := a.analyzeExpression(ret.Value)

			if returnType != lambda.ReturnType && (ret.Value.Kind() != ast.IntLiteralNode && !a.canBeImplicitlyCast(returnType, lambda.ReturnType)) {
				dbg := debug.NewSourceLocationFromExpression(a.Source, ret)
				dbg.ThrowError(
					fmt.Sprintf(
						"Lambda returns '%s', but declared to return '%s'",
						returnType,
						lambda.ReturnType,
					),
					true,
				)
			}
		} else {
			a.analyzeStatement(statement)
		}
	}

	a.Scopes.Pop()

	return ast.NewFunctionType(*lambda)
}
