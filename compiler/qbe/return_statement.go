package qbe

import (
	"blom/ast"
	"strings"
)

func (c Compiler) CompileReturnStatement(stmt *ast.ReturnStatement, ident int) string {
	result := ""

	if stmt.Value.Kind() == ast.IntLiteralNode {
		return "ret " + c.CompileStatement(stmt.Value, ident)
	}

	result += "%r = " + c.CompileStatement(stmt.Value, ident) + "\n"

	result += strings.Repeat("    ", ident-1) + "ret %r"

	return result
}
