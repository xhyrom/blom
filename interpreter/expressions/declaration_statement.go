package expressions

import (
	"blom/ast"
	"blom/env"
)

func InterpretDeclarationStatement(interpreter Interpreter, environment *env.Environment, statement *ast.DeclarationStatement) {
	environment.Set(statement.Name, interpreter.InterpretStatement(statement.Value, environment))
}
