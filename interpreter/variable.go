package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretVariableDeclaration(statement *ast.VariableDeclarationStatement) {
	t.Scopes.Set(statement.Name, t.interpretStatement(statement.Value))
}

func (t *Interpreter) interpretAssignment(statement *ast.Assignment) objects.Object {
	scope, exists := t.Scopes.GetValueScope(statement.Name)
	if !exists {
		panic("missing variable")
	}

	value := t.interpretStatement(statement.Value)

	scope.Set(statement.Name, value)

	return value
}
