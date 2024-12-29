package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (a *TypeAnalyzer) analyzeBlock(block *ast.BlockStatement, scope *env.Environment[*Variable]) compiler.Type {
	newScope := env.New(*scope)

	lastReturnType := compiler.Void

	for _, statement := range block.Body {
		if statement.Kind() == ast.ReturnNode {
			ret := statement.(*ast.ReturnStatement)
			returnType := a.analyzeExpression(ret.Value, newScope)

			handleInconsistentReturnTypes(a, ret, returnType, lastReturnType)
			lastReturnType = returnType
		} else {
			returnType, hasReturnType := a.analyzeStatement(statement, newScope)
			if hasReturnType {
				handleInconsistentReturnTypes(a, statement, returnType, lastReturnType)
				lastReturnType = returnType
			}
		}
	}

	return lastReturnType
}

func handleInconsistentReturnTypes(a *TypeAnalyzer, expression ast.Expression, returnType compiler.Type, lastReturnType compiler.Type) {
	if lastReturnType == compiler.Void || lastReturnType == returnType {
		return
	}

	dbg := debug.NewSourceLocationFromExpression(a.Source, expression)
	dbg.ThrowError(
		fmt.Sprintf(
			"Return type '%s' does not match the previous return type '%s'",
			returnType.Inspect(),
			lastReturnType.Inspect(),
		),
		true,
	)
}
