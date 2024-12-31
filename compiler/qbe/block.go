package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/scope"
	"fmt"
)

func (c *Compiler) compileBlock(block *ast.BlockStatement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	c.Scopes = append(c.Scopes, scope.New[*qbe.TypedValue]())

	c.TempCounter += 1

	blockLabel := fmt.Sprintf("block.%d", c.TempCounter)
	endLabel := fmt.Sprintf("end.%d", c.TempCounter)

	function.AddBlock(blockLabel)

	for _, statement := range block.Body {
		c.compileStatement(statement, function, vtype, isReturn)
	}

	function.AddBlock(endLabel)

	c.Scopes = c.Scopes[:len(c.Scopes)-1]

	return nil
}
