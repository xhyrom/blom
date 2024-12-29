package types

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (a *TypeAnalyzer) analyzeVariableDeclarationStatement(statement *ast.VariableDeclarationStatement, scope *env.Environment[*Variable]) {
	valueType := a.analyzeExpression(statement.Value, scope)

	if statement.Type != valueType {
		dbg := debug.NewSourceLocationFromExpression(a.Source, statement.Value)
		dbg.ThrowError(
			fmt.Sprintf(
				"Variable '%s' declared as '%s', but assigned with '%s'",
				statement.Name,
				statement.Type,
				valueType,
			),
			true,
		)
	}

	scope.Set(statement.Name, &Variable{Type: valueType})
}

func (a *TypeAnalyzer) analyzeAssignmentStatement(statement *ast.AssignmentStatement, scope *env.Environment[*Variable]) {
	variable, exists := scope.FindVariable(statement.Name)
	if !exists {
		dbg := debug.NewSourceLocationFromExpression(a.Source, statement)
		dbg.ThrowError(
			fmt.Sprintf(
				"Variable '%s' is not declared, cannot assign value to it",
				statement.Name,
			),
			true,
			debug.NewHint(
				"Consider declaring the variable before assigning a value to it.",
				"",
			),
		)
	}

	valueType := a.analyzeExpression(statement.Value, scope)

	if variable.Type != valueType {
		dbg := debug.NewSourceLocationFromExpression(a.Source, statement.Value)
		dbg.ThrowError(
			fmt.Sprintf(
				"Variable '%s' declared as '%s', but assigned with '%s'",
				statement.Name,
				variable.Type,
				valueType,
			),
			true,
		)
	}
}
