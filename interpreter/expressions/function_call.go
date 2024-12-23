package expressions

import (
	"blom/ast"
	"blom/env"
)

func InterpretFunctionCall(interpreter Interpreter, environment *env.Environment, call *ast.FunctionCall) env.Object {
	function := environment.FindFunction(call.Name)

	env := env.New(*environment)

	for i, arg := range call.Parameters {
		env.Set(function.Arguments[i].Name, interpreter.InterpretStatement(arg, env))
	}

	return interpreter.InterpretBlock(function.Body, env)
}
