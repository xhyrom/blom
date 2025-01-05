package javascript

import (
	"blom/ast"
)

func (t JavascriptTranspiler) TranspileFunctionDeclaration(declaration *ast.FunctionDeclaration) string {
	result := "function "

	result += declaration.Name + "("

	for i, param := range declaration.Arguments {
		result += param.Name

		if i < len(declaration.Arguments)-1 {
			result += ", "
		}
	}

	result += ") "

	result += t.TranspileBlock(declaration.Body)

	return result
}
