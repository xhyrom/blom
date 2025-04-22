package qbe

import (
	"blom/ast"
	"fmt"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileBlock(block *ast.Block, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
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
