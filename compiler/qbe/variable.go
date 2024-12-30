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

	function.LastBlock().AddInstruction(
		qbe.NewStoreInstruction(variable.Type, value.Value, address.Value),
	)

	return value
}
