package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeIfExpression(expression *ast.IfStatement) compiler.Type {
	conditionType := a.analyzeExpression(expression.Condition)
	if conditionType != compiler.Boolean {
		dbg := debug.NewSourceLocationFromExpression(a.Source, expression.Condition)
		dbg.ThrowError(
			fmt.Sprintf(
				"Condition requires a 'boolean' type, but got '%s'",
				conditionType.Inspect(),
			),
			true,
		)
	}

	a.analyzeStatement(expression.Then)

	if expression.HasElse() {
		a.analyzeStatement(expression.Else)
	}

	return compiler.Void
}
