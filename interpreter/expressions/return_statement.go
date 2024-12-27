package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

func InterpretReturnStatement(interpreter Interpreter, environment *env.Environment, statement *ast.ReturnStatement) objects.Object {
	return interpreter.InterpretStatement(statement.Value, environment)
}
