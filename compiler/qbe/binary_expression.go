package qbe

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) CompileBinaryExpression(stmt *ast.BinaryExpression, indent int) string {
	left := c.CompileStatement(stmt.Left, indent)
	right := c.CompileStatement(stmt.Right, indent)

	result := fmt.Sprintf("%%tmp.%d =%s ", c.Environment.TempCounter, "l")

	switch stmt.Operator {
	case tokens.Plus:
		result += "add"
	}

	result += " " + left + ", " + right

	return result
}
