package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeIf(expression *ast.If) ast.Type {
	conditionType := a.analyzeExpression(expression.Condition)
	if conditionType != ast.Boolean {
		dbg := debug.NewSourceLocationFromExpression(a.Source, expression.Condition)
		dbg.ThrowError(
			fmt.Sprintf(
				"Condition requires a 'boolean' type, but got '%s'",
				conditionType,
			),
			true,
		)
	}

	returnType := a.analyzeBlock(&ast.BlockStatement{
		Body: expression.Then,
		Loc:  expression.Loc,
	})

	if expression.HasElse() {
		elseReturnType := a.analyzeBlock(&ast.BlockStatement{
			Body: expression.Else,
			Loc:  expression.Loc,
		})

		if returnType != elseReturnType {
			dbg := debug.NewSourceLocation(a.Source, expression.Location().Row, expression.Location().Column)
			dbg.ThrowError(
				fmt.Sprintf(
					"Branched blocks must have the same return type, but got '%s' and '%s'",
					returnType,
					elseReturnType,
				),
				true,
			)
		}
	}

	return returnType
}
