package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeWhileLoopStatement(statement *ast.WhileLoopStatement) {
	condition := a.analyzeExpression(statement.Condition)

	if condition != compiler.Boolean {
		dbg := debug.NewSourceLocationFromExpression(a.Source, statement.Condition)
		dbg.ThrowError(
			fmt.Sprintf(
				"Condition requires a 'boolean' type, but got '%s'",
				condition.Inspect(),
			),
			true,
		)
	}

	a.analyzeStatement(statement.Body)
}
