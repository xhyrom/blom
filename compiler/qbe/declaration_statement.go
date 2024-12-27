package qbe

import (
	"blom/ast"
	"blom/env"
	"fmt"
	"strings"
)

func (c *Compiler) CompileDeclarationStatement(stmt *ast.DeclarationStatement, env *env.Environment, indent int) string {
	indentation := strings.Repeat("    ", indent-1)

	result := fmt.Sprintf("%%%s.addr = %s alloc8 4\n", stmt.Name, stmt.Type)
	result += fmt.Sprintf("%sstore%s %s, %%%s.addr\n", indentation, stmt.Type, c.CompileStatement(stmt.Value, env, indent), stmt.Name)
	result += fmt.Sprintf("%s%%%s =%s load%s %%%s.addr", indentation, stmt.Name, stmt.Type, stmt.Type, stmt.Name)

	return result
}
