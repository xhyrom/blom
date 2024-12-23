package javascript

import "blom/ast"

func (t JavascriptTranspiler) TranspileFunctionCall(call *ast.FunctionCall) string {
	result := call.Name + "("

	for i, arg := range call.Parameters {
		result += t.TranspileStatement(arg)

		if i < len(call.Parameters)-1 {
			result += ", "
		}
	}

	result += ")"

	return result
}
