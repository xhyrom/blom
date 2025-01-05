package javascript

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
)

func (t JavascriptTranspiler) TranspileUnaryExpression(expression *ast.UnaryExpression) string {
	operator := ""

	switch expression.Operator {
	case tokens.Plus:
		operator = "+"
	case tokens.Minus:
		operator = "-"
	case tokens.Tilde:
		operator = "~"
	default:
		panic(fmt.Sprintf("unknown operator: %s", expression.Operator))
	}

	return fmt.Sprintf("%s%s", operator, t.TranspileStatement(expression.Operand))
}
