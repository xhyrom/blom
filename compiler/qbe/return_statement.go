package qbe

import (
	"blom/ast"
	"fmt"
)

func (c *Compiler) CompileReturnStatement(stmt *ast.ReturnStatement, ident int) ([]string, string) {
	if c.Environment.CurrentFunction != nil {
		fmt.Printf("Checking return type for function: %s | %T\n", c.Environment.CurrentFunction.Name, stmt.Value)
	}

	result := make([]string, 0)

	s, identifier := c.CompileStatement(stmt.Value, ident)
	for _, s := range s {
		result = append(result, s)
	}

	result = append(result, fmt.Sprintf("ret %s", identifier[2:]))

	result = append(result, "# ^ return statement")

	return result, ""
}
