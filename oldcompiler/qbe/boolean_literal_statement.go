package qbe

import (
	"blom/ast"
	"blom/env"
)

func (c *Compiler) CompileBooleanLiteralStatement(stmt *ast.BooleanLiteral, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	var value int64 = 0

	if stmt.Value {
		value = 1
	}

	return c.CompileStatement(
		&ast.IntLiteral{
			Value: value,
			Type:  stmt.Type,
		},
		scope,
	)
}
