package expressions

import (
	"blom/ast"
	"blom/env"
)

func InterpretReturnStatement(interpreter Interpreter, environment *env.Environment, statement *ast.ReturnStatement) env.Object {
	return interpreter.InterpretExpression(statement.Value, environment)
}
