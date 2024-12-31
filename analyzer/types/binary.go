package types

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

func (a *TypeAnalyzer) analyzeBinaryExpression(expression *ast.BinaryExpression) ast.Type {
	left := a.analyzeExpression(expression.Left)
	right := a.analyzeExpression(expression.Right)

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
