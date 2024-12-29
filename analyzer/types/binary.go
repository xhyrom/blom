package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/env"
	"blom/tokens"
	"fmt"
)

func (a *TypeAnalyzer) analyzeBinaryExpression(expression *ast.BinaryExpression, scope *env.Environment[*Variable]) compiler.Type {
	left := a.analyzeExpression(expression.Left, scope)
	right := a.analyzeExpression(expression.Right, scope)

	if left != right {
		dbg := debug.NewSourceLocation(a.Source, expression.OperatorLoc.Row, expression.OperatorLoc.Column)
		dbg.ThrowError(
			fmt.Sprintf(
				"Binary expression '%s' has mismatched types '%s' and '%s'",
				expression.Operator,
				left.Inspect(),
				right.Inspect(),
			),
			true,
		)
	}

	switch expression.Operator {
	case tokens.Equals:
		return compiler.Boolean
	//case tokens.NotEquals:
	//	return compiler.Boolean
	case tokens.LessThan:
		return compiler.Boolean
	case tokens.LessThanOrEqual:
		return compiler.Boolean
	case tokens.GreaterThan:
		return compiler.Boolean
	case tokens.GreaterThanOrEqual:
		return compiler.Boolean
	}

	return left
}
