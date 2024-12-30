package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileVariableDeclaration(statement *ast.VariableDeclarationStatement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
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
