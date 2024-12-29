package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

func InterpretDeclarationStatement(interpreter Interpreter, environment *env.Environment[objects.Object], statement *ast.VariableDeclarationStatement) {
	obj := interpreter.InterpretStatement(statement.Value, environment)

	environment.Set(statement.Name, obj)
}

func InterpretAssignmentStatement(interpreter Interpreter, environment *env.Environment[objects.Object], statement *ast.AssignmentStatement) {
	_, found := environment.Parent.FindVariable(statement.Name)
	if environment.Parent != nil && found && environment.Get(statement.Name) == nil {
		environment.Parent.Set(statement.Name, interpreter.InterpretStatement(statement.Value, environment))
		return
	}
}
