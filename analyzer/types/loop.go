package types

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (a *TypeAnalyzer) analyzeWhileLoopStatement(statement *ast.WhileLoopStatement, scope *env.Environment[*Variable]) {
	condition := a.analyzeExpression(statement.Condition, scope)

	if condition != ast.Boolean {
		dbg := debug.NewSourceLocationFromExpression(a.Source, statement.Condition)
		dbg.ThrowError(
			fmt.Sprintf(
				"Condition requires a 'boolean' type, but got '%s'",
				condition,
			),
			true,
		)
	}

	a.analyzeStatement(statement.Body, scope)
}
