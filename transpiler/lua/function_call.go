package lua

import "blom/ast"

func (t LuaTranspiler) TranspileFunctionCall(call *ast.FunctionCall) string {
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
