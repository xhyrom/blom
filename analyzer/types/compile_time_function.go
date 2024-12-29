package types

import (
	"blom/ast"
	"blom/debug"
)

func (a *TypeAnalyzer) analyzeCompileTimeFunctionCall(call *ast.CompileTimeFunctionCall) ast.Type {
	switch call.Name {
	case "cast":
		return analyzeCastFunctionCall(a, call)
	default:
		return ast.Null
	}
}

func analyzeCastFunctionCall(a *TypeAnalyzer, call *ast.CompileTimeFunctionCall) ast.Type {
	if len(call.Parameters) != 2 {
		return ast.Null
	}

	literal := call.Parameters[0]
	if literal.Kind() != ast.IdentifierLiteralNode && literal.Kind() != ast.IntLiteralNode && literal.Kind() != ast.FloatLiteralNode {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			"First parameter of 'cast' function must be a literal.",
			true,
		)
	}

	identifierLiteralExpr := call.Parameters[1]
	if identifierLiteralExpr.Kind() != ast.IdentifierLiteralNode {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			"Second parameter of 'cast' function must be an type.",
			true,
		)
	}

	identifierLiteral := identifierLiteralExpr.(*ast.IdentifierLiteral)

	castType, err := ast.ParseType(identifierLiteral.Value)
	if err != nil {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			err.Error(),
			true,
		)
	}

	a.setExpressionType(literal, castType)

	return castType
}
