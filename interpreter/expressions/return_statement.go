package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

func InterpretReturnStatement(interpreter Interpreter, environment *env.Environment[objects.Object], statement *ast.ReturnStatement) objects.Object {
	obj := interpreter.InterpretStatement(statement.Value, environment)

	environment.Collect()

	return obj
}
