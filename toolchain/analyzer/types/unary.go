package types

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

func (a *TypeAnalyzer) analyzeUnaryExpression(expression *ast.UnaryExpression) ast.Type {
	operand := a.analyzeExpression(expression.Operand)

	switch expression.Operator {
	case tokens.Plus:
		expectsType(a, expression, operand, "a numeric", isNumeric)
		return operand
	case tokens.Minus:
		expectsType(a, expression, operand, "a numeric", isNumeric)
		return operand
	case tokens.Tilde:
		expectsType(a, expression, operand, "an integer", isInteger)
		return operand
	case tokens.Ampersand: // address of
		return ast.NewPointerType(operand.AsId())
	case tokens.Asterisk: // dereference
		expectsType(a, expression, operand, "a pointer", isPointer)
		return operand.Dereference()
	}

	return ast.Void
}

type TypeCheckFunc func(ast.Type) bool

func expectsType(a *TypeAnalyzer, expression *ast.UnaryExpression, operand ast.Type, name string, checkFunc TypeCheckFunc) {
	if !checkFunc(operand) {
		dbg := debug.NewSourceLocationFromExpression(a.Source, expression.Operand)
		dbg.ThrowError(
			fmt.Sprintf(
				"Unary expression '%s' expects %s operand, got '%s'",
				expression.Operator,
				name,
				operand,
			),
			true,
		)
	}
}

func isNumeric(t ast.Type) bool {
	return t.IsNumeric()
}

func isInteger(t ast.Type) bool {
	return t.IsInteger()
}

func isPointer(t ast.Type) bool {
	return t.IsPointer()
}
