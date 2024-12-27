package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

func InterpretIfStatement(interpreter Interpreter, environment *env.Environment, statement *ast.IfStatement) objects.Object {
	condition := interpreter.InterpretStatement(statement.Condition, environment)

	if condition == nil {
		return nil
	}

	if condition.(*objects.BooleanObject).Value {
		return interpreter.InterpretStatement(statement.Then, environment)
	}

	if statement.Else != nil {
		return interpreter.InterpretStatement(statement.Else, environment)
	}

	return nil
}
