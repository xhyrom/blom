package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/tokens"
	"fmt"
)

func InterpretBinaryExpression(interpreter Interpreter, environment *env.Environment, expression *ast.BinaryExpression) env.Object {
	return evaluateBinaryExpression(interpreter, environment, expression)
}

func evaluateBinaryExpression(interpreter Interpreter, environment *env.Environment, expression *ast.BinaryExpression) env.Object {
	left := interpreter.InterpretStatement(expression.Left, environment)
	right := interpreter.InterpretStatement(expression.Right, environment)

	switch expression.Operator {
	case tokens.Plus:
		return left.Add(right)
	case tokens.Minus:
		return left.Subtract(right)
	case tokens.Asterisk:
		return left.Multiply(right)
	case tokens.Slash:
		return left.Divide(right)
	default:
		panic(fmt.Sprintf("unknown operator: %s", expression.Operator))
	}
}
