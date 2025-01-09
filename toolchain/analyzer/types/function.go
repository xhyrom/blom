package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
	"strings"
)

func (a *TypeAnalyzer) analyzeFunctionDeclaration(function *ast.FunctionDeclaration) {
	a.Scopes.Append()

	for _, arg := range function.Arguments {
		a.Scopes.Set(arg.Name, &Variable{Type: arg.Type})
	}

	if function.IsNative() {
		return
	}

	for _, statement := range function.Body {
		if statement.Kind() == ast.ReturnNode {
			ret := statement.(*ast.ReturnStatement)
			returnType := a.analyzeExpression(ret.Value)

			if returnType != function.ReturnType && (ret.Value.Kind() != ast.IntLiteralNode && !a.canBeImplicitlyCast(returnType, function.ReturnType)) {
				dbg := debug.NewSourceLocationFromExpression(a.Source, ret)
				dbg.ThrowError(
					fmt.Sprintf(
						"Function '%s' returns '%s', but declared to return '%s'",
						function.Name,
						returnType,
						function.ReturnType,
					),
					true,
				)
			}
		} else {
			a.analyzeStatement(statement)
		}
	}

	function.Name = a.FunctionManager.GetNewName(function)

	a.Scopes.Pop()
}

func (a *TypeAnalyzer) analyzeFunctionCall(call *ast.FunctionCall) ast.Type {
	paramTypes := make([]ast.Type, 0)
	for _, param := range call.Parameters {
		paramTypes = append(paramTypes, a.analyzeExpression(param))
	}

	// TODO: fix the hacky way of changing function call in ast etc...
	function, exists := a.FunctionManager.Get(strings.Split(call.Name, ".")[0], paramTypes)
	if !exists {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)

		overloads := a.FunctionManager.GetAllNamed(call.Name)

		if len(overloads) == 0 {
			dbg.ThrowError(
				fmt.Sprintf("Function '%s' is not defined.", call.PrettyName()),
				true,
			)
		} else if len(overloads) == 1 {
			function = overloads[0]
		} else {
			var overloadStrings []string
			for _, overload := range overloads {
				overloadStrings = append(overloadStrings, formatFunctionSignature(overload))
			}

			message := fmt.Sprintf(
				"No matching overload found for function '%s' with parameter types (%s).\n\n"+
					"Available overloads are:\n%s",
				call.PrettyName(),
				formatTypeList(paramTypes),
				strings.Join(overloadStrings, "\n"),
			)

			dbg.ThrowError(
				message,
				true,
				debug.NewHint(
					"Create an overload with the correct parameter types",
					"",
				),
			)
		}
	}

	if !function.IsNative() && len(function.Arguments) != len(call.Parameters) {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			fmt.Sprintf(
				"Function '%s' (%s) expects %d arguments, but got %d.",
				call.PrettyName(),
				formatFunctionSignature(function),
				len(function.Arguments),
				len(call.Parameters),
			),
			true,
		)
	}

	for i, param := range call.Parameters {
		paramType := a.analyzeExpression(param)

		if !function.IsNative() && paramType != function.Arguments[i].Type && !a.canBeImplicitlyCast(paramType, function.Arguments[i].Type) {
			dbg := debug.NewSourceLocation(a.Source, param.Location().Row, param.Location().Column)
			dbg.ThrowError(
				fmt.Sprintf(
					"Function '%s' (%s) expects argument %d to be of type '%s', but got '%s'.",
					call.PrettyName(),
					formatFunctionSignature(function),
					i+1,
					function.Arguments[i].Type,
					paramType,
				),
				true,
			)
		}
	}

	if !function.HasAnnotation(ast.Infix) && call.Infix {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			fmt.Sprintf("Function '%s' is not marked as infix.", call.PrettyName()),
			true,
			debug.NewHint(
				fmt.Sprintf(
					"Mark the function '%s' at %d:%d as infix",
					call.PrettyName(),
					function.Location().Row,
					function.Location().Column,
				),
				"",
			),
		)
	}

	call.Name = a.FunctionManager.GetNewName(function)

	return function.ReturnType
}

func formatTypeList(types []ast.Type) string {
	typeNames := make([]string, len(types))
	for i, t := range types {
		typeNames[i] = t.String()
	}

	return strings.Join(typeNames, ", ")
}

func formatFunctionSignature(fn *ast.FunctionDeclaration) string {
	return fmt.Sprintf("  %s(%s)", fn.PrettyName(), formatTypeList(getArgumentTypes(fn)))
}

func getArgumentTypes(fn *ast.FunctionDeclaration) []ast.Type {
	types := make([]ast.Type, len(fn.Arguments))
	for i, param := range fn.Arguments {
		types[i] = param.Type
	}
	return types
}
