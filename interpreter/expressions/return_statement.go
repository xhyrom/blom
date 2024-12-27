package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"blom/env/objects"
	"fmt"
)

func InterpretReturnStatement(interpreter Interpreter, environment *env.Environment, statement *ast.ReturnStatement) objects.Object {
	obj := interpreter.InterpretStatement(statement.Value, environment)
	currentFunction := environment.FindCurrentFunction()

	if currentFunction != nil && currentFunction.ReturnType != obj.Type() {
		dbg := debug.NewSourceLocation(interpreter.Source(), statement.Loc.Row, statement.Loc.Column)
		dbg.ThrowError(fmt.Sprintf("Expected return type %s, got %s", currentFunction.ReturnType, obj.Type()), true)
	}

	environment.Collect()

	return obj
}
