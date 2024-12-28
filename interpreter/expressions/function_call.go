package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"blom/env/objects"
	"fmt"
)

func InterpretFunctionCall(interpreter Interpreter, environment *env.Environment[objects.Object], call *ast.FunctionCall) objects.Object {
	function := environment.FindFunction(call.Name)

	env := env.New(*environment)

	for i, param := range call.Parameters {
		arg := function.Arguments[i]
		obj := interpreter.InterpretStatement(param, env)

		if arg.Type != obj.Type() {
			dbg := debug.NewSourceLocation(interpreter.Source(), call.Loc.Row, call.Loc.Column)
			dbg.ThrowError(fmt.Sprintf("Expected type %s, got %s", arg.Type, obj.Type()), true)
		}

		env.Set(function.Arguments[i].Name, obj)
	}

	env.CurrentFunction = function

	return interpreter.InterpretBlock(function.Body, env)
}
