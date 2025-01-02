package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeVariableDeclarationStatement(statement *ast.VariableDeclarationStatement) {
	valueType := a.analyzeExpression(statement.Value)

	if statement.Type != valueType && !a.canBeImplicitlyCast(valueType, statement.Type) {
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

func (a *TypeAnalyzer) analyzeAssignment(assignment *ast.Assignment) ast.Type {
	variable := a.getVariable(assignment.Name)
	if variable == nil {
		dbg := debug.NewSourceLocationFromExpression(a.Source, assignment)
		dbg.ThrowError(
			fmt.Sprintf(
				"Variable '%s' is not declared, cannot assign value to it",
				assignment.Name,
			),
			true,
			debug.NewHint(
				"Consider declaring the variable before assigning a value to it.",
				"",
			),
		)
	}

	valueType := a.analyzeExpression(assignment.Value)

	if variable.Type != valueType && !a.canBeImplicitlyCast(variable.Type, valueType) {
		dbg := debug.NewSourceLocationFromExpression(a.Source, assignment.Value)
		dbg.ThrowError(
			fmt.Sprintf(
				"Variable '%s' declared as '%s', but assigned with '%s'",
				assignment.Name,
				variable.Type,
				valueType,
			),
			true,
		)
	}

	return valueType
}
