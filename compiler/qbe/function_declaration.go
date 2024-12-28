package qbe

import (
	"blom/ast"
)

func (c *Compiler) CompileFunctionDeclaration(stmt *ast.FunctionDeclaration, ident int) []string {
	c.Environment.SetFunction(stmt.Name, stmt)

	if stmt.IsNative() {
		return []string{}
	}

	c.Environment.CurrentFunction = stmt

	result := ""

	if stmt.Name == "main" {
		result += "export "
	}

	result += "function " + stmt.ReturnType.String() + " $" + stmt.Name + "("

	for i, param := range stmt.Arguments {
		if i > 0 {
			result += ", "
		}

		result += param.Type.String() + " %" + param.Name
	}

	result += ") {\n"
	result += "@start\n"

	block := c.CompileBlock(*stmt.Body, ident)
	for _, b := range block {
		result += b
	}

	result += "}"

	return []string{result, "# ^ function declaration\n"}
}
