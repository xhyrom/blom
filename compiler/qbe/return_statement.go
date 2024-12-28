package qbe

import (
	"blom/ast"
	"blom/compiler"
	"fmt"
)

func (c *Compiler) CompileReturnStatement(stmt *ast.ReturnStatement, ident int, expectedType *compiler.Type) ([]string, *Additional) {
	if c.Environment.CurrentFunction != nil {
		fmt.Printf("Checking return type for function: %s | %T\n", c.Environment.CurrentFunction.Name, stmt.Value)
	}

	result := make([]string, 0)

	s, identifier := c.CompileStatement(stmt.Value, ident, expectedType)
	for _, s := range s {
		result = append(result, s)
	}

	result = append(result, fmt.Sprintf("ret %s", identifier.Name))

	result = append(result, "# ^ return statement")

	return result, nil
}
