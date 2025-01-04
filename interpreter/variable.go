package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretVariableDeclaration(statement *ast.VariableDeclarationStatement, function *ast.FunctionDeclaration, isReturn bool) {
	ty := statement.Type

	value := t.interpretStatement(statement.Value, function, &ty, isReturn)

	if ty != value.Type() {
		cnv := t.convertToType(value.Type(), ty, value)
		value = cnv
	}

	t.Scopes.Set(statement.Name, value)
}

func (t *Interpreter) interpretAssignment(statement *ast.Assignment, function *ast.FunctionDeclaration, isReturn bool) objects.Object {
	scope, exists := t.Scopes.GetValueScope(statement.Name)
	if !exists {
		panic("missing variable")
	}

	variable, _ := scope.Get(statement.Name)

	ty := variable.Type()
	value := t.interpretStatement(statement.Value, function, &ty, isReturn)

	if ty != value.Type() {
		cnv := t.convertToType(value.Type(), ty, value)
		value = cnv
	}

	scope.Set(statement.Name, value)

	return value
}
