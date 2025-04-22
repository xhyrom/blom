package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Codegen) compileBlock(block *ast.Block, function *qbe.Function, vtype qbe.Type, isReturn bool) *qbe.TypedValue {
	c.Scopes.Append()

	c.TempCounter += 1

	blockLabel := fmt.Sprintf("block.%d.body", c.TempCounter)
	endLabel := fmt.Sprintf("block.%d.end", c.TempCounter)

	function.AddBlock(blockLabel)

	for _, statement := range block.Body {
		c.compileStatement(statement, function, vtype, isReturn)
	}

	function.AddBlock(endLabel)

	c.Scopes.Pop()

	return nil
}
