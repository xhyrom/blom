package lua

import (
	"blom/ast"
)

func (t LuaTranspiler) TranspileFunctionDeclaration(declaration *ast.FunctionDeclaration) string {
	result := "function "

	result += declaration.Name + "("

	for i, param := range declaration.Arguments {
		result += param.Name

		if i < len(declaration.Arguments)-1 {
			result += ", "
		}
	}

	result += ")\n"

	body := t.TranspileBlock(declaration.Body)
	body = body[3:] // remove the "do" from the beginning

	result += body

	return result
}
