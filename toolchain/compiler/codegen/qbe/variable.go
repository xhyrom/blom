package qbe

import (
	"blom/ast"
	"blom/tokens"
	"fmt"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileVariableDeclaration(statement *ast.VariableDeclaration, function *ir.Function, isReturn bool) *ir.TypedValue {
	t := RemapAstType(statement.Type)

	value := c.compileStatement(statement.Init, function, t, isReturn)
	if value.Type.IsFunction() {
		t = value.Type
	}

	c.createVariable(t, statement.Name.Value)
	address := c.createVariable(t, fmt.Sprintf("%s.addr", statement.Name.Value))

	function.LastBlock().AddAssign(
		address,
		ir.NewPointer(t),
		ir.Alloc8Instruction{
			Value: ir.NewConstantValue(int64(t.Size(c.Module))),
		},
	)

	if !value.Type.IsFunction() && t != value.Type {
		value = c.convertToType(value.Type, t, value.Value, function)
		t = value.Type
	}

	function.LastBlock().AddInstruction(
		ir.NewStoreInstruction(t.IntoBase(), value.Value, address),
	)

	return value
}

func (c *Codegen) compileAssignmentStatement(statement *ast.Assignment, function *ir.Function, isReturn bool) *ir.TypedValue {
	address := evaluateLeftSide(c, statement.Left, function)
	value := c.compileStatement(statement.Right, function, address.Type, isReturn)

	t := address.Type
	if !value.Type.IsFunction() && t != value.Type {
		value = c.convertToType(value.Type, t, value.Value, function)
		t = value.Type
	}

	function.LastBlock().AddInstruction(
		ir.NewStoreInstruction(t, value.Value, address.Value),
	)

	return value
}

func evaluateLeftSide(c *Codegen, left ast.Node, function *ir.Function) *ir.TypedValue {
	switch expr := left.(type) {
	case *ast.IdentifierLiteral:
		address, exists := c.Scopes.GetValue(fmt.Sprintf("%s.addr", expr.Value))
		if !exists {
			panic("missing variable address")
		}

		return address

	case *ast.UnaryExpression:
		if expr.Operator != tokens.Asterisk {
			panic("unsupported unary operator")
		}

		operand := c.compileStatement(expr.Operand, function, nil, false)
		return operand

	default:
		panic("unsupported left expression")
	}
}
