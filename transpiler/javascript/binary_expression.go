package javascript

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
)

func (t JavascriptTranspiler) TranspileBinaryExpression(expression *ast.BinaryExpression) string {
	left := t.TranspileStatement(expression.Left)
	right := t.TranspileStatement(expression.Right)

	operator := ""

	switch expression.Operator {
	case tokens.Plus:
		operator = "+"
	case tokens.Minus:
		operator = "-"
	case tokens.Asterisk:
		operator = "*"
	case tokens.Slash:
		operator = "/"
	case tokens.PercentSign:
		operator = "%"
	case tokens.Ampersand:
		operator = "&"
	case tokens.VerticalLine:
		operator = "|"
	case tokens.CircumflexAccent:
		operator = "~"
	case tokens.DoubleLessThan:
		operator = "<<"
	case tokens.DoubleGreaterThan:
		operator = ">>"
	case tokens.Equals:
		operator = "=="
	case tokens.LessThan:
		operator = "<"
	case tokens.LessThanOrEqual:
		operator = "<="
	case tokens.GreaterThan:
		operator = ">"
	case tokens.GreaterThanOrEqual:
		operator = ">="
	default:
		panic(fmt.Sprintf("unknown operator: %s", expression.Operator))
	}

	if isBinaryExpression(expression.Left) && expression.Operator.Precedence() > expression.Left.(*ast.BinaryExpression).Operator.Precedence() {
		left = fmt.Sprintf("(%s)", left)
	}
	if isBinaryExpression(expression.Right) && expression.Operator.Precedence() > expression.Right.(*ast.BinaryExpression).Operator.Precedence() {
		right = fmt.Sprintf("(%s)", right)
	}

	return fmt.Sprintf("%s %s %s", left, operator, right)
}

func isBinaryExpression(statement ast.Statement) bool {
	_, ok := statement.(*ast.BinaryExpression)
	return ok
}
