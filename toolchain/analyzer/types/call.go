package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
	"strings"
)

func (a *TypeAnalyzer) analyzeCall(call ast.Call) ast.Type {
	switch call := call.(type) {
	case *ast.FunctionCall:
		return a.analyzeFunctionCall(call)
	case *ast.MethodCall:
		return a.analyzeMethodCall(call)
	case *ast.InfixCall:
		return a.analyzeInfixCall(call)
	}

	return nil
}

func (a *TypeAnalyzer) analyzeFunctionCall(call *ast.FunctionCall) ast.Type {
	name := call.Path.Dotify()

	argTypes := make([]ast.Type, 0)
	for _, arg := range call.Args {
		argTypes = append(argTypes, a.analyzeExpression(arg))
	}

	function, _ := a.FunctionManager.Get(name, argTypes)
	if function == nil {
		return ast.Void
	}

	a.Context.FunctionCalls[call] = function

	isNative := function.HasAnnotation(ast.Native)

	if !isNative && len(function.Params) != len(call.Args) {
		dbg := debug.NewSourceLocationFromNode(a.Source, call)
		dbg.ThrowError(
			fmt.Sprintf(
				"Function '%s' (%s) expects %d arguments, but got %d.",
				name,
				formatFunctionSignature(function),
				len(function.Params),
				len(call.Args),
			),
			true,
		)
	}

	for i, arg := range call.Args {
		argType := argTypes[i]

		if !isNative && !argType.Equal(function.Params[i].Type) && !a.canBeImplicitlyCast(argType, function.Params[i].Type) {
			dbg := debug.NewSourceLocation(a.Source, arg.Location().Row, arg.Location().Column)
			dbg.ThrowError(
				fmt.Sprintf(
					"Function '%s' (%s) expects argument %d to be of type '%s', but got '%s'.",
					name,
					formatFunctionSignature(function),
					i+1,
					function.Params[i].Type,
					argType,
				),
				true,
			)
		}
	}

	return function.Return
}

func (a *TypeAnalyzer) analyzeMethodCall(call *ast.MethodCall) ast.Type {
	// TODO: implement method call
	panic("not implemented method call")
}

func (a *TypeAnalyzer) analyzeInfixCall(call *ast.InfixCall) ast.Type {
	if call.FunctionCall != nil {
		return a.analyzeFunctionCall(call.FunctionCall)
	}

	return a.analyzeMethodCall(call.MethodCall)
}

func formatTypeList(types []ast.Type) string {
	typeNames := make([]string, len(types))
	for i, t := range types {
		typeNames[i] = t.String()
	}

	return strings.Join(typeNames, ", ")
}

func formatFunctionSignature(fn *ast.FunctionDeclaration) string {
	return fmt.Sprintf("  %s(%s)", fn.Name, formatTypeList(getArgumentTypes(fn)))
}

func getArgumentTypes(fn *ast.FunctionDeclaration) []ast.Type {
	types := make([]ast.Type, len(fn.Params))
	for i, param := range fn.Params {
		types[i] = param.Type
	}
	return types
}
