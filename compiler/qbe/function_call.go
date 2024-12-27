package qbe

import "blom/ast"

func (c Compiler) CompileFunctionCall(stmt *ast.FunctionCall, ident int) string {
	result := "call $" + stmt.Name + "("

	for i, param := range stmt.Parameters {
		if i > 0 {
			result += ", "
		}

		result += c.CompileStatement(param, ident)
	}

	result += ")"

	return result
}
