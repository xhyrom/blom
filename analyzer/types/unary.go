package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/env"
	"blom/tokens"
	"fmt"
)

func (a *TypeAnalyzer) analyzeUnaryExpression(expression *ast.UnaryExpression, scope *env.Environment[*Variable]) compiler.Type {
	operand := a.analyzeExpression(expression.Operand, scope)

	if !operand.IsNumeric() {
		dbg := debug.NewSourceLocationFromExpression(a.Source, expression.Operand)
		dbg.ThrowError(
			fmt.Sprintf(
				"Unary expression '%s' expects a numeric operand, got '%s'",
				expression.Operator,
				operand.Inspect(),
			),
			true,
		)
	}

	switch expression.Operator {
	case tokens.Plus:
		return operand
	case tokens.Minus:
		return operand
	case tokens.Tilde:
		if !operand.IsInteger() {
			dbg := debug.NewSourceLocationFromExpression(a.Source, expression.Operand)
			dbg.ThrowError(
				fmt.Sprintf(
					"Unary expression '~' expects an integer operand, got '%s'",
					operand.Inspect(),
				),
				true,
			)
		}

		return operand
	}

	return compiler.Void
}
