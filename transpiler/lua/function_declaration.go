package lua

import (
	"blom/ast"
	"strings"
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

	bodyLines := strings.Split(body, "\n")
	for _, line := range bodyLines {
		result += "   " + line + "\n"
	}

	result += "end\n"

	return result
}
