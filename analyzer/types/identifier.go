package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeIdentifier(expression *ast.IdentifierLiteralStatement) compiler.Type {
	variable := a.Environment.Get(expression.Value)
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
