package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) compileUnaryExpression(expression *ast.UnaryExpression, function *qbe.Function, vtype qbe.Type, isReturn bool) *qbe.TypedValue {
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
	case tokens.Ampersand: // address of
		return compileAddressOf(c, expression.Operand, function, vtype)
	case tokens.Asterisk: // dereference
		return compileDereference(c, expression.Operand, function, vtype)
	}

	panic(fmt.Sprintf("unknown unary operator: %s", expression.Operator))
}

func compileAddressOf(c *Compiler, expression ast.Expression, function *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	val := c.compileStatement(expression, function, vtype, false)
	ty := qbe.NewPointer(val.Type)

	if val := expression.(*ast.IdentifierLiteral); val != nil {
		if _, exists := c.Scopes.GetValue(val.Value); exists {
			address, exists := c.Scopes.GetValue(fmt.Sprintf("%s.addr", val.Value))

			if exists {
				return &qbe.TypedValue{
					Type:  ty,
					Value: address.Value,
				}
			}
		}
	}

	tempValue := c.getTemporaryValue(nil)

	function.LastBlock().AddAssign(
		tempValue,
		ty,
		qbe.NewAlloc8Instruction(qbe.NewConstantValue(ty.Size())),
	)

	function.LastBlock().AddInstruction(
		qbe.NewStoreInstruction(
			ty,
			tempValue,
			val.Value,
		),
	)

	return &qbe.TypedValue{
		Type:  ty,
		Value: tempValue,
	}
}

func compileDereference(c *Compiler, expression ast.Expression, function *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	val := c.compileStatement(expression, function, vtype, false)
	tempValue := c.getTemporaryValue(nil)

	function.LastBlock().AddAssign(
		tempValue,
		val.Type.(qbe.PointerBox).Inner,
		qbe.NewLoadInstruction(
			val.Type.(qbe.PointerBox).Inner,
			val.Value,
		),
	)

	return &qbe.TypedValue{
		Type:  val.Type.(qbe.PointerBox).Inner,
		Value: tempValue,
	}
}
