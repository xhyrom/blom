package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeWhileLoopStatement(statement *ast.WhileLoopStatement) {
	condition := a.analyzeExpression(statement.Condition)

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

	a.analyzeStatement(&ast.BlockStatement{
		Body: statement.Body,
		Loc:  statement.Loc,
	})
}
