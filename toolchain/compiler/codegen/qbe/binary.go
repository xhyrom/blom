package qbe

import (
	"blom/ast"
	"blom/tokens"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileBinaryExpression(expression *ast.BinaryExpression, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
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

	var instruction ir.Instruction
	switch expression.Operator {
	case tokens.Plus:
		instruction = ir.NewAddInstruction(left, right)
	case tokens.Minus:
		instruction = ir.NewSubtractInstruction(left, right)
	case tokens.Asterisk:
		instruction = ir.NewMultiplyInstruction(left, right)
	case tokens.Slash:
		instruction = ir.NewDivideInstruction(left, right)
	case tokens.PercentSign:
		instruction = ir.NewModulusInstruction(left, right)
	case tokens.LessThan:
		instruction = ir.NewCompareInstruction(
			ty,
			ir.LessThan,
			left,
			right,
		)
	case tokens.LessThanOrEqual:
		instruction = ir.NewCompareInstruction(
			ty,
			ir.LessThanOrEqual,
			left,
			right,
		)
	case tokens.GreaterThan:
		instruction = ir.NewCompareInstruction(
			ty,
			ir.GreaterThan,
			left,
			right,
		)
	case tokens.GreaterThanOrEqual:
		instruction = ir.NewCompareInstruction(
			ty,
			ir.GreaterThanOrEqual,
			left,
			right,
		)
	case tokens.Equals:
		instruction = ir.NewCompareInstruction(
			ty,
			ir.Equal,
			left,
			right,
		)
	case tokens.NotEquals:
		instruction = ir.NewCompareInstruction(
			ty,
			ir.NotEqual,
			left,
			right,
		)
	case tokens.Ampersand:
		instruction = ir.NewBitwiseAndInstruction(left, right)
	case tokens.VerticalLine:
		instruction = ir.NewBitwiseOrInstruction(left, right)
	case tokens.CircumflexAccent:
		instruction = ir.NewBitwiseXorInstruction(left, right)
	case tokens.DoubleLessThan:
		instruction = ir.NewShiftLeftInstruction(left, right)
	case tokens.DoubleGreaterThan:
		instruction = ir.NewArithmeticShiftRightInstruction(left, right)
	}

	tempValue := c.getTemporaryValue(nil)

	if isComparisonOperator(expression.Operator) {
		ty = ir.Boolean
	}

	function.LastBlock().AddAssign(
		tempValue,
		ty,
		instruction,
	)

	return &ir.TypedValue{
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
