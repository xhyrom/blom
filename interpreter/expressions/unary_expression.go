package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
	"blom/tokens"
)

func InterpretUnaryExpression(interpreter Interpreter, environment *env.Scope[objects.Object], expression *ast.UnaryExpression) objects.Object {
	switch expression.Operator {
	case tokens.Plus:
		return evaluatePlusExpression(interpreter, environment, expression)
	case tokens.Minus:
		return evaluateMinusExpression(interpreter, environment, expression)
	case tokens.Tilde:
		return evaluateBitwiseNotExpression(interpreter, environment, expression)
	}
	return nil
}

func evaluatePlusExpression(interpreter Interpreter, environment *env.Scope[objects.Object], expression *ast.UnaryExpression) objects.Object {
	return interpreter.InterpretStatement(expression.Operand, environment)
}

func evaluateMinusExpression(interpreter Interpreter, environment *env.Scope[objects.Object], expression *ast.UnaryExpression) objects.Object {
	operand := interpreter.InterpretStatement(expression.Operand, environment)

	switch operand := operand.(type) {
	case *objects.DoubleObject:
		return &objects.DoubleObject{Value: -operand.Value}
	case *objects.FloatObject:
		return &objects.FloatObject{Value: -operand.Value}
	case *objects.LongObject:
		return &objects.LongObject{Value: -operand.Value}
	case *objects.IntObject:
		return &objects.IntObject{Value: -operand.Value}
	case *objects.ShortObject:
		return &objects.ShortObject{Value: -operand.Value}
	case *objects.ByteObject:
		return &objects.ByteObject{Value: -operand.Value}
	}

	return nil
}

func evaluateBitwiseNotExpression(interpreter Interpreter, environment *env.Scope[objects.Object], expression *ast.UnaryExpression) objects.Object {
	operand := interpreter.InterpretStatement(expression.Operand, environment)

	switch operand := operand.(type) {
	case *objects.LongObject:
		return &objects.LongObject{Value: ^operand.Value}
	case *objects.IntObject:
		return &objects.IntObject{Value: ^operand.Value}
	case *objects.ShortObject:
		return &objects.ShortObject{Value: ^operand.Value}
	case *objects.ByteObject:
		return &objects.ByteObject{Value: ^operand.Value}
	}

	return nil
}
