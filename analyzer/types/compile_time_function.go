package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/env"
)

func (a *TypeAnalyzer) analyzeCompileTimeFunctionCall(call *ast.CompileTimeFunctionCall, scope *env.Environment[*Variable]) compiler.Type {
	switch call.Name {
	case "cast":
		return analyzeCastFunctionCall(a, call, scope)
	default:
		return compiler.Null
	}
}

func analyzeCastFunctionCall(a *TypeAnalyzer, call *ast.CompileTimeFunctionCall, scope *env.Environment[*Variable]) compiler.Type {
	if len(call.Parameters) != 2 {
		return compiler.Null
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

	identifierLiteral := identifierLiteralExpr.(*ast.IdentifierLiteralStatement)

	castType, err := compiler.ParseType(identifierLiteral.Value)
	if err != nil {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			err.Error(),
			true,
		)
	}

	a.setExpressionType(literal, castType, scope)

	return castType
}
