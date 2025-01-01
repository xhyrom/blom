package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileVariableDeclaration(statement *ast.VariableDeclarationStatement, function *qbe.Function, isReturn bool) *qbe.TypedValue {
	t := qbe.RemapAstType(statement.Type)

	value := c.compileStatement(statement.Value, function, &t, isReturn)

	c.createVariable(t, statement.Name)
	address := c.createVariable(t, fmt.Sprintf("%s.addr", statement.Name))

	function.LastBlock().AddAssign(
		address,
		qbe.NewPointer(t),
		qbe.Alloc8Instruction{
			Value: qbe.NewConstantValue(int64(t.Size())),
		},
	)

	if t != value.Type {
		cnv := c.convertToType(value.Type, t, value.Value, function)
		t = cnv.Type
	}

	function.LastBlock().AddInstruction(
		qbe.NewStoreInstruction(t, value.Value, address),
	)

	return value
}

func (c *Compiler) compileAssignmentStatement(statement *ast.AssignmentStatement, function *qbe.Function, isReturn bool) *qbe.TypedValue {
	variable := c.getVariable(statement.Name)
	if variable == nil {
		panic("missing variable")
	}

	value := c.compileStatement(statement.Value, function, &variable.Type, isReturn)

	address := c.getVariable(fmt.Sprintf("%s.addr", statement.Name))
	if address == nil {
		panic("missing address")
	}

	t := variable.Type

	if t != value.Type {
		cnv := c.convertToType(value.Type, t, value.Value, function)
		t = cnv.Type
		value.Type = t
	}

	function.LastBlock().AddInstruction(
		qbe.NewStoreInstruction(t, value.Value, address.Value),
	)

	return value
}
