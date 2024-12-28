package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"fmt"
)

func (c *Compiler) CompileReturnStatement(stmt *ast.ReturnStatement, ident int, expectedType *compiler.Type) ([]string, *Additional) {
	result := make([]string, 0)

	s, identifier := c.CompileStatement(stmt.Value, ident, expectedType)
	for _, s := range s {
		result = append(result, s)
	}

	if c.Environment.CurrentFunction != nil && c.Environment.CurrentFunction.ReturnType != identifier.Type {
		dbg := debug.NewSourceLocation(c.Source, stmt.Loc.Row, stmt.Loc.Column)
		dbg.ThrowError(fmt.Sprintf("Return type mismatch: expected \"%s\", got \"%s\"", c.Environment.CurrentFunction.ReturnType.Inspect(), identifier.Type.Inspect()), true)
	}

	result = append(result, fmt.Sprintf("ret %s", identifier.Name))

	result = append(result, "# ^ return statement")

	return result, nil
}
