package qbe

import (
	"blom/ast"
	"blom/tokens"
)

func (c Compiler) CompileBinaryExpression(stmt *ast.BinaryExpression, ident int) string {
	left := c.CompileStatement(stmt.Left, ident)
	right := c.CompileStatement(stmt.Right, ident)

	result := "w "

	switch stmt.Operator {
	case tokens.Plus:
		result += "add"
	}

	result += " " + left + ", " + right

	return result
}
