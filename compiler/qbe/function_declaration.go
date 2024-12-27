package qbe

import "blom/ast"

func (c Compiler) CompileFunctionDeclaration(stmt *ast.FunctionDeclaration, ident int) string {
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

	result += c.CompileBlock(*stmt.Body, ident)

	result += "}\n"

	return result
}
