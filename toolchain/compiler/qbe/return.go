package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileReturnStatement(statement *ast.ReturnStatement, function *qbe.Function, vtype *qbe.Type) *qbe.TypedValue {
	returnStatement := c.compileStatement(statement.Value, function, vtype, true)

	if returnStatement == nil {
		function.LastBlock().AddInstruction(qbe.ReturnInstruction{
			Value: nil,
		})
	} else {
		function.LastBlock().AddInstruction(qbe.ReturnInstruction{
			Value: returnStatement.Value,
		})
	}

	return nil
}
