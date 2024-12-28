package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

func InterpretWhileLoopStatement(interpreter Interpreter, environment *env.Environment[objects.Object], statement *ast.WhileLoopStatement) {
	condition := statement.Condition
	body := statement.Body

	for {
		conditionValue := interpreter.InterpretStatement(condition, environment)
		if conditionValue == nil {
			break
		}

		if conditionValue.(*objects.BooleanObject).Value {
			interpreter.InterpretBlock(body, environment)
		} else {
			break
		}
	}
}
