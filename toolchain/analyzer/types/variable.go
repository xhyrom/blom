package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeVariableDeclarationStatement(statement *ast.VariableDeclarationStatement) {
	valueType := a.analyzeExpression(statement.Value)

	if statement.Type != valueType && (!a.canBeImplicitlyCast(valueType, statement.Type) || statement.Value.Kind() == ast.FunctionCallNode) {
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

	if valueType.IsPointer() && valueType.Dereference() == ast.Void {
		lambda := statement.Value.(*ast.LambdaDeclaration)

		a.FunctionManager.Register(&ast.FunctionDeclaration{
			Name:       statement.Name,
			Arguments:  lambda.Arguments,
			ReturnType: lambda.ReturnType,
			Body:       lambda.Body,
			Loc:        statement.Loc,
		})
	} else {
		a.Scopes.Set(statement.Name, &Variable{Type: valueType})
	}
}

func (a *TypeAnalyzer) analyzeAssignment(assignment *ast.Assignment) ast.Type {
	leftType := a.analyzeExpression(assignment.Left)
	rightType := a.analyzeExpression(assignment.Right)

	if leftType != rightType && !a.canBeImplicitlyCast(rightType, leftType) {
		dbg := debug.NewSourceLocationFromExpression(a.Source, assignment.Right)
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
