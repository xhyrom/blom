package types

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"blom/tokens"
	"fmt"
)

func (a *TypeAnalyzer) analyzeBinaryExpression(expression *ast.BinaryExpression, scope *env.Environment[*Variable]) ast.Type {
	left := a.analyzeExpression(expression.Left, scope)
	right := a.analyzeExpression(expression.Right, scope)

	if left != right {
		dbg := debug.NewSourceLocation(a.Source, expression.OperatorLoc.Row, expression.OperatorLoc.Column)
		dbg.ThrowError(
			fmt.Sprintf(
				"Binary expression '%s' has mismatched types '%s' and '%s'",
				expression.Operator,
				left,
				right,
			),
			true,
		)
	}

	switch expression.Operator {
	case tokens.Equals:
		return ast.Boolean
	//case tokens.NotEquals:
	//	return compiler.Boolean
	case tokens.LessThan:
		return ast.Boolean
	case tokens.LessThanOrEqual:
		return ast.Boolean
	case tokens.GreaterThan:
		return ast.Boolean
	case tokens.GreaterThanOrEqual:
		return ast.Boolean
	}

	return left
}
