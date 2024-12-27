package qbe

import (
	"blom/ast"
	"blom/env"
	"blom/tokens"
)

func (c *Compiler) CompileBinaryExpression(stmt *ast.BinaryExpression, env *env.Environment, ident int) string {
	left := c.CompileStatement(stmt.Left, env, ident)
	right := c.CompileStatement(stmt.Right, env, ident)

	result := "w "

	switch stmt.Operator {
	case tokens.Plus:
		result += "add"
	}

	result += " " + left + ", " + right

	return result
}
