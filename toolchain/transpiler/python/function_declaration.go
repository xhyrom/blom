package python

import (
	"blom/ast"
	"blom/env"
)

func (t PythonTranspiler) TranspileFunctionDeclaration(declaration *ast.FunctionDeclaration, environment *env.Scope, indent int) string {
	result := "def "

	result += declaration.Name + "("

	newEnv := env.New(*environment)

	for i, param := range declaration.Arguments {
		result += param.Name

		if i < len(declaration.Arguments)-1 {
			result += ", "
		}

		newEnv.Set(declaration.Name, &env.BooleanObject{
			Value: false,
		})
	}

	result += "):\n"

	result += t.TranspileBlock(declaration.Body, newEnv, indent+1)[9:] // removes "if True:\n"

	return result
}
