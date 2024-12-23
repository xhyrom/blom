package expressions

import (
	"blom/ast"
	"blom/env"
)

func InterpretDeclarationStatement(interpreter Interpreter, environment *env.Environment, statement *ast.DeclarationStatement) {
	if statement.Redeclaration {
		if environment.Parent != nil && environment.Parent.FindVariable(statement.Name) != nil && environment.Get(statement.Name) == nil {
			environment.Parent.Set(statement.Name, interpreter.InterpretStatement(statement.Value, environment))
			return
		}
	}

	environment.Set(statement.Name, interpreter.InterpretStatement(statement.Value, environment))
}
