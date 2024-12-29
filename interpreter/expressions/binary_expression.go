package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
	"blom/tokens"
	"fmt"
)

func InterpretBinaryExpression(interpreter Interpreter, environment *env.Environment[objects.Object], expression *ast.BinaryExpression) objects.Object {
	return evaluateBinaryExpression(interpreter, environment, expression)
}

func evaluateBinaryExpression(interpreter Interpreter, environment *env.Environment[objects.Object], expression *ast.BinaryExpression) objects.Object {
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
	case tokens.PercentSign:
		return left.Modulo(right)
	case tokens.Ampersand:
		return left.BitwiseAnd(right)
	case tokens.VerticalLine:
		return left.BitwiseOr(right)
	case tokens.CircumflexAccent:
		return left.BitwiseXor(right)
	case tokens.DoubleLessThan:
		return left.LeftShift(right)
	case tokens.DoubleGreaterThan:
		return left.RightShift(right)
	case tokens.Equals:
		return left.Equals(right)
	case tokens.LessThan:
		return left.LessThan(right)
	case tokens.LessThanOrEqual:
		return left.LessThanOrEqual(right)
	case tokens.GreaterThan:
		return left.GreaterThan(right)
	case tokens.GreaterThanOrEqual:
		return left.GreaterThanOrEqual(right)
	default:
		panic(fmt.Sprintf("unknown operator: %s", expression.Operator))
	}
}
