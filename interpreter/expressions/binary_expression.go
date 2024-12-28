package expressions

import (
	"blom/ast"
	"blom/debug"
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
		return handleTypeMismatch(interpreter, expression, left, right, left.Add(right))
	case tokens.Minus:
		return handleTypeMismatch(interpreter, expression, left, right, left.Subtract(right))
	case tokens.Asterisk:
		return handleTypeMismatch(interpreter, expression, left, right, left.Multiply(right))
	case tokens.Slash:
		return handleTypeMismatch(interpreter, expression, left, right, left.Divide(right))
	case tokens.PercentSign:
		return handleTypeMismatch(interpreter, expression, left, right, left.Modulo(right))
	case tokens.Ampersand:
		return handleTypeMismatch(interpreter, expression, left, right, left.BitwiseAnd(right))
	case tokens.VerticalLine:
		return handleTypeMismatch(interpreter, expression, left, right, left.BitwiseOr(right))
	case tokens.CircumflexAccent:
		return handleTypeMismatch(interpreter, expression, left, right, left.BitwiseXor(right))
	case tokens.DoubleLessThan:
		return handleTypeMismatch(interpreter, expression, left, right, left.LeftShift(right))
	case tokens.DoubleGreaterThan:
		return handleTypeMismatch(interpreter, expression, left, right, left.RightShift(right))
	case tokens.Equals:
		return handleTypeMismatch(interpreter, expression, left, right, left.Equals(right))
	case tokens.LessThan:
		return handleTypeMismatch(interpreter, expression, left, right, left.LessThan(right))
	case tokens.LessThanOrEqual:
		return handleTypeMismatch(interpreter, expression, left, right, left.LessThanOrEqual(right))
	case tokens.GreaterThan:
		return handleTypeMismatch(interpreter, expression, left, right, left.GreaterThan(right))
	case tokens.GreaterThanOrEqual:
		return handleTypeMismatch(interpreter, expression, left, right, left.GreaterThanOrEqual(right))
	default:
		panic(fmt.Sprintf("unknown operator: %s", expression.Operator))
	}
}

func handleTypeMismatch(interpreter Interpreter, expression *ast.BinaryExpression, left objects.Object, right objects.Object, object objects.Object) objects.Object {
	if object == nil {
		dbg := debug.NewSourceLocation(interpreter.Source(), expression.OperatorLoc.Row, expression.OperatorLoc.Column)
		dbg.ThrowError(fmt.Sprintf("Type mismatch: %s %s %s", left.Type(), expression.Operator, right.Type()), true)
	}

	return object
}
