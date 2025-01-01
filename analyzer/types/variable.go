package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeVariableDeclarationStatement(statement *ast.VariableDeclarationStatement) {
	valueType := a.analyzeExpression(statement.Value)

	if statement.Type != valueType && !a.canBeImplicitlyCast(statement.Type, valueType) {
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

	a.createVariable(statement.Name, &Variable{Type: valueType})
}

func (a *TypeAnalyzer) analyzeAssignmentStatement(statement *ast.AssignmentStatement) {
	variable := a.getVariable(statement.Name)
	if variable == nil {
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

	valueType := a.analyzeExpression(statement.Value)

	if variable.Type != valueType && !a.canBeImplicitlyCast(variable.Type, valueType) {
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
