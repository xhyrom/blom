package types

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (a *TypeAnalyzer) analyzeIdentifier(expression *ast.IdentifierLiteral, scope *env.Environment[*Variable]) ast.Type {
	variable, exists := scope.FindVariable(expression.Value)

	if !exists {
		dbg := debug.NewSourceLocationFromExpression(a.Source, expression)
		dbg.ThrowError(
			fmt.Sprintf(
				"Variable '%s' is not declared",
				expression.Value,
			),
			true,
		)
	}

	return variable.Type
}
