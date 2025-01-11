package types

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

func (a *TypeAnalyzer) analyzeUnaryExpression(expression *ast.UnaryExpression) ast.Type {
	operand := a.analyzeExpression(expression.Operand)

	if !operand.IsNumeric() {
		dbg := debug.NewSourceLocationFromExpression(a.Source, expression.Operand)
		dbg.ThrowError(
			fmt.Sprintf(
				"Unary expression '%s' expects a numeric operand, got '%s'",
				expression.Operator,
				operand,
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
					operand,
				),
				true,
			)
		}

		return operand
	case tokens.Ampersand: // address of
		// todo: implement
		return operand
	}

	return ast.Void
}
