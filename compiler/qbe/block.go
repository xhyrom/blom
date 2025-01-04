package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileBlock(block *ast.BlockStatement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	c.Scopes.Append()

	c.TempCounter += 1

	blockLabel := fmt.Sprintf("block.%d", c.TempCounter)
	endLabel := fmt.Sprintf("end.%d", c.TempCounter)

	function.AddBlock(blockLabel)

	for _, statement := range block.Body {
		c.compileStatement(statement, function, vtype, isReturn)
	}

	function.AddBlock(endLabel)

	c.Scopes.Pop()

	return nil
}
