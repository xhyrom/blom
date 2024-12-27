package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"fmt"
)

func InterpretDeclarationStatement(interpreter Interpreter, environment *env.Environment, statement *ast.DeclarationStatement) {
	if statement.Redeclaration {
		if environment.Parent != nil && environment.Parent.FindVariable(statement.Name) != nil && environment.Get(statement.Name) == nil {
			environment.Parent.Set(statement.Name, interpreter.InterpretStatement(statement.Value, environment))
			return
		}
	}

	obj := interpreter.InterpretStatement(statement.Value, environment)

	if statement.Type != obj.Type() {
		dbg := debug.NewSourceLocation(interpreter.Source(), statement.Location().Row, statement.Location().Column)
		dbg.ThrowError(fmt.Sprintf("Type mismatch in declaration: %s != %s", statement.Type.Inspect(), obj.Type().Inspect()), true)
	}

	environment.Set(statement.Name, obj)
}
