package qbe

import (
	"blom/ast"
	"blom/env"
	"strings"
)

func (c *Compiler) CompileReturnStatement(stmt *ast.ReturnStatement, env *env.Environment, ident int) string {
	result := ""

	if stmt.Value.Kind() == ast.IntLiteralNode {
		return "ret " + c.CompileStatement(stmt.Value, env, ident)
	}

	result += "%r = " + c.CompileStatement(stmt.Value, env, ident) + "\n"

	result += strings.Repeat("    ", ident-1) + "ret %r"

	return result
}
