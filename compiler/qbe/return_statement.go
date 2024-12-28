package qbe

import (
	"blom/ast"
	"fmt"
	"strings"
)

func (c *Compiler) CompileReturnStatement(stmt *ast.ReturnStatement, ident int) string {
	if c.Environment.CurrentFunction != nil {
		fmt.Printf("Checking return type for function: %s | %T\n", c.Environment.CurrentFunction.Name, stmt.Value)
	}

	result := ""

	if stmt.Value.Kind() == ast.IntLiteralNode {
		return "ret " + c.CompileStatement(stmt.Value, ident)
	}

	if stmt.Value.Kind() == ast.StringLiteralNode {
		s := c.CompileStatement(stmt.Value, ident)
		return "ret " + s[2:]
	}

	result += fmt.Sprintf("%s\n", c.CompileStatement(stmt.Value, ident))
	result += fmt.Sprintf("%sret %%tmp.%d", strings.Repeat("    ", ident-1), c.Environment.TempCounter)

	c.Environment.TempCounter += 1

	return result
}
