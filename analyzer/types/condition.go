package types

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (a *TypeAnalyzer) analyzeIfStatement(expression *ast.IfStatement, scope *env.Environment[*Variable]) {
	conditionType := a.analyzeExpression(expression.Condition, scope)
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

	a.analyzeBlock(&ast.BlockStatement{
		Body: expression.Then,
		Loc:  expression.Loc,
	}, scope)

	if expression.HasElse() {
		a.analyzeBlock(&ast.BlockStatement{
			Body: expression.Else,
			Loc:  expression.Loc,
		}, scope)
	}
}

func (a *TypeAnalyzer) analyzeIfExpression(expression *ast.IfStatement, scope *env.Environment[*Variable]) ast.Type {
	conditionType := a.analyzeExpression(expression.Condition, scope)
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
	}, scope)

	if expression.HasElse() {
		elseReturnType := a.analyzeBlock(&ast.BlockStatement{
			Body: expression.Else,
			Loc:  expression.Loc,
		}, scope)

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
