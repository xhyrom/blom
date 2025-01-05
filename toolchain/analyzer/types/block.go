package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeBlock(block *ast.BlockStatement) ast.Type {
	a.Scopes.Append()

	lastReturnType := ast.Void

	for _, statement := range block.Body {
		if statement.Kind() == ast.ReturnNode {
			ret := statement.(*ast.ReturnStatement)
			returnType := a.analyzeExpression(ret.Value)

			//handleInconsistentReturnTypes(a, ret, returnType, lastReturnType)
			lastReturnType = returnType
		} else {
			returnType, hasReturnType := a.analyzeStatement(statement)
			if hasReturnType {
				//handleInconsistentReturnTypes(a, statement, returnType, lastReturnType)
				lastReturnType = returnType
			}
		}
	}

	a.Scopes.Pop()

	return lastReturnType
}

func handleInconsistentReturnTypes(a *TypeAnalyzer, expression ast.Expression, returnType ast.Type, lastReturnType ast.Type) {
	if lastReturnType == ast.Void || lastReturnType == returnType {
		return
	}

	dbg := debug.NewSourceLocationFromExpression(a.Source, expression)
	dbg.ThrowError(
		fmt.Sprintf(
			"Return type '%s' does not match the previous return type '%s'",
			returnType,
			lastReturnType,
		),
		true,
	)
}
