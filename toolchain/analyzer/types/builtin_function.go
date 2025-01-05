package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeBuiltinFunctionCall(call *ast.BuiltinFunctionCall) ast.Type {
	switch call.Name {
	case "cast":
		return analyzeCastFunctionCall(a, call)
	}

	dbg := debug.NewSourceLocationFromExpression(a.Source, call)
	dbg.ThrowError(
		fmt.Sprintf("Unknown builtin function call '%s'.", call.Name),
		true,
	)

	return ast.Null
}

func analyzeCastFunctionCall(a *TypeAnalyzer, call *ast.BuiltinFunctionCall) ast.Type {
	if len(call.Parameters) != 2 {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			"Function 'cast' requires exactly two parameters (type, expression).",
			true,
		)
	}

	firstParam := call.Parameters[0]
	if firstParam.Kind() != ast.IdentifierLiteralNode {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			"First parameter of 'cast' function must be an type.",
			true,
		)
	}

	typeName := firstParam.(*ast.IdentifierLiteral).Value
	castType, err := ast.ParseType(typeName)
	if err != nil {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			err.Error(),
			true,
		)
	}

	ret := a.analyzeExpression(call.Parameters[1])
	if ret == ast.Void {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			"Expression in 'cast' function must return a value.",
			true,
		)
	}

	return castType
}
