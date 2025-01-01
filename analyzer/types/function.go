package types

import (
	"blom/ast"
	"blom/debug"
	"blom/scope"
	"fmt"
)

func (a *TypeAnalyzer) analyzeFunctionDeclaration(function *ast.FunctionDeclaration) {
	a.Scopes = append(a.Scopes, scope.New[*Variable]())

	for _, arg := range function.Arguments {
		a.Scopes[len(a.Scopes)-1].Set(arg.Name, &Variable{Type: arg.Type})
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
}

func (a *TypeAnalyzer) analyzeFunctionCall(call *ast.FunctionCall) ast.Type {
	function, exists := a.Functions.Get(call.Name)
	if !exists {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			fmt.Sprintf(
				"Function '%s' is not defined.",
				call.Name,
			),
			true,
		)
	}

	if !function.IsNative() && len(function.Arguments) != len(call.Parameters) {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			fmt.Sprintf(
				"Function '%s' expects %d arguments, but got %d.",
				call.Name,
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
					"Function '%s' expects argument %d to be of type '%s', but got '%s'.",
					call.Name,
					i+1,
					function.Arguments[i].Type,
					paramType,
				),
				true,
			)
		}
	}

	return function.ReturnType
}
