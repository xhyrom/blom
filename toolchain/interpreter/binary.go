package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
	"blom/tokens"
)

func (t *Interpreter) interpretBinaryExpression(expression *ast.BinaryExpression, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	typedLeft := t.interpretStatement(expression.Left, function, vtype, isReturn)
	typedRight := t.interpretStatement(expression.Right, function, vtype, isReturn)

	leftType := typedLeft.Type()
	left := typedLeft

	rightType := typedRight.Type()
	right := typedRight

	if leftType.Weight() > rightType.Weight() {
		typedRight = t.convertToType(rightType, leftType, right)

		rightType = typedRight.Type()
		right = typedRight
	} else if leftType.Weight() < rightType.Weight() {
		typedLeft = t.convertToType(leftType, rightType, left)

		leftType = typedLeft.Type()
		left = typedLeft
	}

	ty := typedLeft.Type()

	var obj objects.Object
	switch expression.Operator {
	case tokens.Plus:
		obj = left.Add(right)
	case tokens.Minus:
		obj = left.Subtract(right)
	case tokens.Asterisk:
		obj = left.Multiply(right)
	case tokens.Slash:
		obj = left.Divide(right)
	case tokens.PercentSign:
		obj = left.Modulo(right)
	case tokens.LessThan:
		obj = left.LessThan(right)
	case tokens.LessThanOrEqual:
		obj = left.LessThanOrEqual(right)
	case tokens.GreaterThan:
		obj = left.GreaterThan(right)
	case tokens.GreaterThanOrEqual:
		obj = left.GreaterThanOrEqual(right)
	case tokens.Equals:
		obj = left.Equals(right)
	case tokens.Ampersand:
		obj = left.BitwiseAnd(right)
	case tokens.VerticalLine:
		obj = left.BitwiseOr(right)
	case tokens.CircumflexAccent:
		obj = left.BitwiseXor(right)
	case tokens.DoubleLessThan:
		obj = left.LeftShift(right)
	case tokens.DoubleGreaterThan:
		obj = left.RightShift(right)
	}

	if isComparisonOperator(expression.Operator) {
		ty = ast.Boolean
	}

	newObj := objects.FromType(ty)
	newObj.SetValue(obj.Value())

	return newObj
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
