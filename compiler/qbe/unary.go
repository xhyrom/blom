package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) compileUnaryExpression(expression *ast.UnaryExpression, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	switch expression.Operator {
	case tokens.Plus: // unary plus
		return c.compileStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: 1,
			},
			Loc:         expression.Operand.Location(),
			Operator:    tokens.Asterisk, // multiply by 1
			OperatorLoc: expression.Operand.Location(),
		}, function, vtype, isReturn)
	case tokens.Minus: // unary minus
		return c.compileStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: -1,
			},
			Loc:         expression.Operand.Location(),
			Operator:    tokens.Asterisk, // multiply by -1
			OperatorLoc: expression.Operand.Location(),
		}, function, vtype, isReturn)
	case tokens.Tilde: // bitwise not
		return c.compileStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: -1,
			},
			Loc:         expression.Operand.Location(),
			Operator:    tokens.CircumflexAccent, // bitwise xor
			OperatorLoc: expression.Operand.Location(),
		}, function, vtype, isReturn)
	}

	panic(fmt.Sprintf("unknown unary operator: %s", expression.Operator))
}
