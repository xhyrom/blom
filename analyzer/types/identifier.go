package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeIdentifier(expression *ast.IdentifierLiteral) ast.Type {
	variable := a.getVariable(expression.Value)

	if variable == nil {
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
