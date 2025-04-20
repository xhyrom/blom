package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeVariableDeclarationStatement(statement *ast.VariableDeclaration) {
	valueType := a.analyzeExpression(statement.Init)

	if !statement.Type.Equal(valueType) && (!a.canBeImplicitlyCast(valueType, statement.Type) || statement.Init.Kind() == ast.FunctionCallNode) {
		dbg := debug.NewSourceLocationFromNode(a.Source, statement.Init)
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

	a.Scopes.Set(statement.Name.Value, &Variable{Type: valueType})
}

func (a *TypeAnalyzer) analyzeAssignment(assignment *ast.Assignment) ast.Type {
	leftType := a.analyzeExpression(assignment.Left)
	rightType := a.analyzeExpression(assignment.Right)

	if !leftType.Equal(rightType) && !a.canBeImplicitlyCast(rightType, leftType) {
		dbg := debug.NewSourceLocationFromNode(a.Source, assignment.Right)
		dbg.ThrowError(
			fmt.Sprintf(
				"Cannot assign value of type '%s' to '%s'",
				rightType,
				leftType,
			),
			true,
		)
	}

	return leftType
}
