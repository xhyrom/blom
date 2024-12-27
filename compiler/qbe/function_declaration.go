package qbe

import (
	"blom/ast"
	"blom/env"
)

func (c *Compiler) CompileFunctionDeclaration(stmt *ast.FunctionDeclaration, env *env.Environment, ident int) string {
	env.SetFunction(stmt.Name, stmt)

	if stmt.IsNative() {
		return ""
	}

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

	result += c.CompileBlock(*stmt.Body, env, ident)

	result += "}\n"

	return result
}
