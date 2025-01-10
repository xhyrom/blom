package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/tokens"
)

func (c *Compiler) compileBinaryExpression(expression *ast.BinaryExpression, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	typedLeft := c.compileStatement(expression.Left, function, vtype, isReturn)
	typedRight := c.compileStatement(expression.Right, function, vtype, isReturn)

	leftType := typedLeft.Type
	left := typedLeft.Value

	rightType := typedRight.Type
	right := typedRight.Value

	if leftType.Weight() > rightType.Weight() {
		typedRight = c.convertToType(rightType, leftType, right, function)

		rightType = typedRight.Type
		right = typedRight.Value
	} else if leftType.Weight() < rightType.Weight() {
		typedLeft = c.convertToType(leftType, rightType, left, function)

		leftType = typedLeft.Type
		left = typedLeft.Value
	}

	ty := typedLeft.Type

	var instruction qbe.Instruction
	switch expression.Operator {
	case tokens.Plus:
		instruction = qbe.NewAddInstruction(left, right)
	case tokens.Minus:
		instruction = qbe.NewSubtractInstruction(left, right)
	case tokens.Asterisk:
		instruction = qbe.NewMultiplyInstruction(left, right)
	case tokens.Slash:
		instruction = qbe.NewDivideInstruction(left, right)
	case tokens.PercentSign:
		instruction = qbe.NewModulusInstruction(left, right)
	case tokens.LessThan:
		instruction = qbe.NewCompareInstruction(
			ty,
			qbe.LessThan,
			left,
			right,
		)
	case tokens.LessThanOrEqual:
		instruction = qbe.NewCompareInstruction(
			ty,
			qbe.LessThanOrEqual,
			left,
			right,
		)
	case tokens.GreaterThan:
		instruction = qbe.NewCompareInstruction(
			ty,
			qbe.GreaterThan,
			left,
			right,
		)
	case tokens.GreaterThanOrEqual:
		instruction = qbe.NewCompareInstruction(
			ty,
			qbe.GreaterThanOrEqual,
			left,
			right,
		)
	case tokens.Equals:
		instruction = qbe.NewCompareInstruction(
			ty,
			qbe.Equal,
			left,
			right,
		)
	case tokens.NotEquals:
		instruction = qbe.NewCompareInstruction(
			ty,
			qbe.NotEqual,
			left,
			right,
		)
	case tokens.Ampersand:
		instruction = qbe.NewBitwiseAndInstruction(left, right)
	case tokens.VerticalLine:
		instruction = qbe.NewBitwiseOrInstruction(left, right)
	case tokens.CircumflexAccent:
		instruction = qbe.NewBitwiseXorInstruction(left, right)
	case tokens.DoubleLessThan:
		instruction = qbe.NewShiftLeftInstruction(left, right)
	case tokens.DoubleGreaterThan:
		instruction = qbe.NewArithmeticShiftRightInstruction(left, right)
	}

	tempValue := c.getTemporaryValue(nil)

	if isComparisonOperator(expression.Operator) {
		ty = qbe.Boolean
	}

	function.LastBlock().AddAssign(
		tempValue,
		ty,
		instruction,
	)

	return &qbe.TypedValue{
		Type:  ty,
		Value: tempValue,
	}
}

func isComparisonOperator(operator tokens.TokenKind) bool {
	switch operator {
	case tokens.LessThan,
		tokens.LessThanOrEqual,
		tokens.GreaterThan,
		tokens.GreaterThanOrEqual,
		tokens.Equals:
		return true
	}

	return false
}
