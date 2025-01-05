package python

import (
	"blom/ast"
	"blom/env"
	"blom/tokens"
	"fmt"
)

func (t PythonTranspiler) TranspileUnaryExpression(expression *ast.UnaryExpression, environment *env.Scope, indent int) string {
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

	return fmt.Sprintf("%s%s", operator, t.TranspileStatement(expression.Operand, environment, indent))
}
