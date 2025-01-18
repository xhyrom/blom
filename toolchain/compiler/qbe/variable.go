package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) compileVariableDeclaration(statement *ast.VariableDeclarationStatement, function *qbe.Function, isReturn bool) *qbe.TypedValue {
	t := qbe.RemapAstType(statement.Type)

	value := c.compileStatement(statement.Value, function, &t, isReturn)
	if value.Type.IsFunction() {
		t = value.Type
	}

	c.createVariable(t, statement.Name)
	address := c.createVariable(t, fmt.Sprintf("%s.addr", statement.Name))

	function.LastBlock().AddAssign(
		address,
		qbe.NewPointer(t),
		qbe.Alloc8Instruction{
			Value: qbe.NewConstantValue(int64(t.Size())),
		},
	)

	if !value.Type.IsFunction() && t != value.Type {
		value = c.convertToType(value.Type, t, value.Value, function)
		t = value.Type
	}

	function.LastBlock().AddInstruction(
		qbe.NewStoreInstruction(t, value.Value, address),
	)

	return value
}

func (c *Compiler) compileAssignmentStatement(statement *ast.Assignment, function *qbe.Function, isReturn bool) *qbe.TypedValue {
	address := evaluateLeftSide(c, statement.Left, function)

	value := c.compileStatement(statement.Right, function, &address.Type, isReturn)

	t := address.Type
	if !value.Type.IsFunction() && t != value.Type {
		value = c.convertToType(value.Type, t, value.Value, function)
		t = value.Type
	}

	function.LastBlock().AddInstruction(
		qbe.NewStoreInstruction(t, value.Value, address.Value),
	)

	return value
}

func evaluateLeftSide(c *Compiler, left ast.Expression, function *qbe.Function) *qbe.TypedValue {
	switch expr := left.(type) {
	case *ast.IdentifierLiteral:
		address, exists := c.Scopes.GetValue(fmt.Sprintf("%s.addr", expr.Value))
		if !exists {
			panic("missing address")
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
