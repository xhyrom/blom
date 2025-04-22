package qbe

import (
	"blom/ast"
	"blom/tokens"
	"fmt"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileUnaryExpression(expression *ast.UnaryExpression, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
	switch expression.Operator {
	case tokens.Plus: // unary plus
		return c.compileStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: 1,
			},
			Loc:      expression.Operand.Location(),
			Operator: tokens.Asterisk, // multiply by 1
		}, function, vtype, isReturn)
	case tokens.Minus: // unary minus
		return c.compileStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: -1,
			},
			Loc:      expression.Operand.Location(),
			Operator: tokens.Asterisk, // multiply by -1
		}, function, vtype, isReturn)
	case tokens.Tilde: // bitwise not
		return c.compileStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: -1,
			},
			Loc:      expression.Operand.Location(),
			Operator: tokens.CircumflexAccent, // bitwise xor
		}, function, vtype, isReturn)
	case tokens.Ampersand: // address of
		return compileAddressOf(c, expression.Operand, function, vtype)
	case tokens.Asterisk: // dereference
		return compileDereference(c, expression.Operand, function, vtype)
	}

	panic(fmt.Sprintf("unknown unary operator: %s", expression.Operator))
}

func compileAddressOf(c *Codegen, expression ast.Node, function *ir.Function, vtype ir.Type) *ir.TypedValue {
	val := c.compileStatement(expression, function, vtype, false)
	ty := ir.NewPointer(val.Type)

	if val := expression.(*ast.IdentifierLiteral); val != nil {
		if _, exists := c.Scopes.GetValue(val.Value); exists {
			address, exists := c.Scopes.GetValue(fmt.Sprintf("%s.addr", val.Value))

			if exists {
				return &ir.TypedValue{
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
		ir.NewAlloc8Instruction(ir.NewConstantValue(ty.Size(c.Module))),
	)

	function.LastBlock().AddInstruction(
		ir.NewStoreInstruction(
			ty,
			tempValue,
			val.Value,
		),
	)

	return &ir.TypedValue{
		Type:  ty,
		Value: tempValue,
	}
}

func compileDereference(c *Codegen, expression ast.Node, function *ir.Function, vtype ir.Type) *ir.TypedValue {
	val := c.compileStatement(expression, function, vtype, false)
	tempValue := c.getTemporaryValue(nil)

	function.LastBlock().AddAssign(
		tempValue,
		val.Type.(ir.PointerBox).Inner.IntoBase(),
		ir.NewLoadInstruction(
			val.Type.(ir.PointerBox).Inner.IntoBase(),
			val.Value,
		),
	)

	return &ir.TypedValue{
		Type:  val.Type.(ir.PointerBox).Inner,
		Value: tempValue,
	}
}
