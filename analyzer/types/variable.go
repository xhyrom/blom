package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeVariableDeclarationStatement(statement *ast.VariableDeclarationStatement) {
	valueType := a.analyzeExpression(statement.Value)

	if statement.Type != valueType {
		dbg := debug.NewSourceLocationFromExpression(a.Source, statement.Value)
		dbg.ThrowError(
			fmt.Sprintf(
				"Variable '%s' declared as '%s', but assigned with '%s'",
				statement.Name,
				statement.Type.Inspect(),
				valueType.Inspect(),
			),
			true,
		)
	}

	a.Environment.Set(statement.Name, &Variable{Type: valueType})
}

func (a *TypeAnalyzer) analyzeAssignmentStatement(statement *ast.AssignmentStatement) {
	variable := a.Environment.Get(statement.Name)
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

	if variable.Type != valueType {
		dbg := debug.NewSourceLocationFromExpression(a.Source, statement.Value)
		dbg.ThrowError(
			fmt.Sprintf(
				"Variable '%s' declared as '%s', but assigned with '%s'",
				statement.Name,
				variable.Type.Inspect(),
				valueType.Inspect(),
			),
			true,
		)
	}
}
