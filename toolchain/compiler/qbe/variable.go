package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) compileVariableDeclaration(statement *ast.VariableDeclaration, function *qbe.Function, isReturn bool) *qbe.TypedValue {
	t := qbe.RemapAstType(statement.Type)

	value := c.compileStatement(statement.Init, function, t, isReturn)
	if value.Type.IsFunction() {
		t = value.Type
	}

	c.createVariable(t, statement.Name.Value)
	address := c.createVariable(t, fmt.Sprintf("%s.addr", statement.Name.Value))

	function.LastBlock().AddAssign(
		address,
		qbe.NewPointer(t),
		qbe.Alloc8Instruction{
			Value: qbe.NewConstantValue(int64(t.Size(c.Module))),
		},
	)

	if !value.Type.IsFunction() && t != value.Type {
		value = c.convertToType(value.Type, t, value.Value, function)
		t = value.Type
	}

	function.LastBlock().AddInstruction(
		qbe.NewStoreInstruction(t.IntoBase(), value.Value, address),
	)

	return value
}

func (c *Compiler) compileAssignmentStatement(statement *ast.Assignment, function *qbe.Function, isReturn bool) *qbe.TypedValue {
	address := evaluateLeftSide(c, statement.Left, function)
	value := c.compileStatement(statement.Right, function, address.Type, isReturn)

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

func evaluateLeftSide(c *Compiler, left ast.Node, function *qbe.Function) *qbe.TypedValue {
	switch expr := left.(type) {
	case *ast.IdentifierLiteral:
		address, exists := c.Scopes.GetValue(fmt.Sprintf("%s.addr", expr.Value))
		if !exists {
			val, _ := c.Scopes.GetValue(expr.Value)
			address := c.createVariable(val.Type, fmt.Sprintf("%s.addr", expr.Value))

			function.FirstBlock().AddAssignAt(
				0,
				address,
				qbe.NewPointer(val.Type),
				qbe.Alloc8Instruction{
					Value: qbe.NewConstantValue(int64(val.Type.Size(c.Module))),
				},
			)

			return evaluateLeftSide(c, expr, function)
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
