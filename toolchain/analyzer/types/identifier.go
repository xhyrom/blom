package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeIdentifier(expression *ast.IdentifierLiteral) ast.Type {
	variable, exists := a.Scopes.GetValue(expression.Value)

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
