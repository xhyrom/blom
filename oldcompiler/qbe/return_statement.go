package qbe

import (
	"blom/ast"
	"blom/env"
	"fmt"
)

func (c *Compiler) CompileReturnStatement(stmt *ast.ReturnStatement, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	result := make([]string, 0)

	s, identifier := c.CompileStatement(stmt.Value, scope)
	for _, s := range s {
		result = append(result, s)
	}

	result = append(result, fmt.Sprintf("ret %s", identifier.Name))
	result = append(result, "# ^ return statement")

	return result, nil
}
