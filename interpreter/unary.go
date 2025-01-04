package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
	"blom/tokens"
	"fmt"
)

func (t *Interpreter) interpretUnaryExpression(expression *ast.UnaryExpression, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	switch expression.Operator {
	case tokens.Plus: // unary plus
		return t.interpretStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: 1,
			},
			Loc:         expression.Operand.Location(),
			Operator:    tokens.Asterisk, // multiply by 1
			OperatorLoc: expression.Operand.Location(),
		}, function, vtype, isReturn)
	case tokens.Minus: // unary minus
		return t.interpretStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: -1,
			},
			Loc:         expression.Operand.Location(),
			Operator:    tokens.Asterisk, // multiply by -1
			OperatorLoc: expression.Operand.Location(),
		}, function, vtype, isReturn)
	case tokens.Tilde: // bitwise not
		return t.interpretStatement(&ast.BinaryExpression{
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
