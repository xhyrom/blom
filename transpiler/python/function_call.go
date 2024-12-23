package python

import (
	"blom/ast"
	"blom/env"
)

func (t PythonTranspiler) TranspileFunctionCall(call *ast.FunctionCall, environment *env.Environment, indent int) string {
	result := call.Name + "("

	for i, arg := range call.Parameters {
		result += t.TranspileStatement(arg, environment, indent)

		if i < len(call.Parameters)-1 {
			result += ", "
		}
	}

	result += ")"

	return result
}
