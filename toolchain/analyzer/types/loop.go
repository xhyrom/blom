package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeWhileLoopStatement(statement *ast.WhileLoop) {
	condition := a.analyzeExpression(statement.Condition)

	if condition != ast.Boolean {
		dbg := debug.NewSourceLocationFromNode(a.Source, statement.Condition)
		dbg.ThrowError(
			fmt.Sprintf(
				"Condition requires a 'boolean' type, but got '%s'",
				condition,
			),
			true,
		)
	}

	a.analyzeStatement(&ast.Block{
		Body: statement.Body,
		Loc:  statement.Loc,
	})
}
