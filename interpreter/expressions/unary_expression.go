package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/tokens"
)

func InterpretUnaryExpression(interpreter Interpreter, environment *env.Environment, expression *ast.UnaryExpression) env.Object {
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

func evaluatePlusExpression(interpreter Interpreter, environment *env.Environment, expression *ast.UnaryExpression) env.Object {
	return interpreter.InterpretStatement(expression.Operand, environment)
}

func evaluateMinusExpression(interpreter Interpreter, environment *env.Environment, expression *ast.UnaryExpression) env.Object {
	operand := interpreter.InterpretStatement(expression.Operand, environment)

	switch operand := operand.(type) {
	case *env.IntegerObject:
		return &env.IntegerObject{Value: -operand.Value}
	case *env.FloatObject:
		return &env.FloatObject{Value: -operand.Value}
	}

	return nil
}

func evaluateBitwiseNotExpression(interpreter Interpreter, environment *env.Environment, expression *ast.UnaryExpression) env.Object {
	operand := interpreter.InterpretStatement(expression.Operand, environment)

	switch operand := operand.(type) {
	case *env.IntegerObject:
		return &env.IntegerObject{Value: ^operand.Value}
	}

	return nil
}
