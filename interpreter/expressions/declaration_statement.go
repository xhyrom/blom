package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"blom/env/objects"
	"fmt"
)

func InterpretDeclarationStatement(interpreter Interpreter, environment *env.Environment[objects.Object], statement *ast.DeclarationStatement) {
	if statement.Redeclaration {
		_, found := environment.Parent.FindVariable(statement.Name)
		if environment.Parent != nil && found && environment.Get(statement.Name) == nil {
			environment.Parent.Set(statement.Name, interpreter.InterpretStatement(statement.Value, environment))
			return
		}
	}

	obj := interpreter.InterpretStatement(statement.Value, environment)

	if *statement.Type != obj.Type() {
		dbg := debug.NewSourceLocation(interpreter.Source(), statement.Location().Row, statement.Location().Column)
		dbg.ThrowError(fmt.Sprintf("Type mismatch in declaration: %s != %s", statement.Type.Inspect(), obj.Type().Inspect()), true)
	}

	environment.Set(statement.Name, obj)
}
