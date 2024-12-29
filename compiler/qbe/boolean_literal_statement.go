package qbe

import (
	"blom/ast"
	"blom/compiler"
)

func (c *Compiler) CompileBooleanLiteralStatement(stmt *ast.BooleanLiteralStatement, expectedType *compiler.Type) ([]string, *Additional) {
	var value int64 = 0

	if stmt.Value {
		value = 1
	}

	return c.CompileStatement(
		&ast.IntLiteralStatement{
			Value: value,
		},
		expectedType,
	)
}
