package qbe

import (
	"blom/ast"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileReturnStatement(statement *ast.Return, function *ir.Function, vtype ir.Type) *ir.TypedValue {
	returnStatement := c.compileStatement(statement.Value, function, vtype, true)

	if returnStatement == nil {
		function.LastBlock().AddInstruction(ir.ReturnInstruction{
			Value: nil,
		})
	} else {
		function.LastBlock().AddInstruction(ir.ReturnInstruction{
			Value: returnStatement.Value,
		})
	}

	return nil
}
