package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (a *TypeAnalyzer) analyzeIfExpression(expression *ast.IfStatement, scope *env.Environment[*Variable]) compiler.Type {
	conditionType := a.analyzeExpression(expression.Condition, scope)
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

	a.analyzeStatement(expression.Then, scope)

	if expression.HasElse() {
		a.analyzeStatement(expression.Else, scope)
	}

	return compiler.Void
}
