package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (a *TypeAnalyzer) analyzeFunctionDeclaration(function *ast.FunctionDeclaration, scope *env.Environment[*Variable]) {
	functionScope := env.New[*Variable](*scope)

	for _, arg := range function.Arguments {
		functionScope.Set(arg.Name, &Variable{Type: arg.Type})
	}

	if function.IsNative() {
		return
	}

	for _, statement := range function.Body.Body {
		if statement.Kind() == ast.ReturnNode {
			ret := statement.(*ast.ReturnStatement)
			returnType := a.analyzeExpression(ret.Value, functionScope)

			if returnType != function.ReturnType {
				dbg := debug.NewSourceLocationFromExpression(a.Source, ret)
				dbg.ThrowError(
					fmt.Sprintf(
						"Function '%s' returns '%s', but declared to return '%s'",
						function.Name,
						returnType.Inspect(),
						function.ReturnType.Inspect(),
					),
					true,
				)
			}
		} else {
			a.analyzeStatement(statement, functionScope)
		}
	}
}

func (a *TypeAnalyzer) analyzeFunctionCall(call *ast.FunctionCall, scope *env.Environment[*Variable]) compiler.Type {
	function := scope.FindFunction(call.Name)
	if function == nil {
		dbg := debug.NewSourceLocationFromExpression(a.Source, call)
		dbg.ThrowError(
			fmt.Sprintf(
				"Function '%s' is not defined.",
				call.Name,
			),
			true,
		)
	}

	if function.IsNative() {
		return function.ReturnType
	}

	if len(function.Arguments) != len(call.Parameters) {
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
		paramType := a.analyzeExpression(param, scope)

		if paramType != function.Arguments[i].Type {
			dbg := debug.NewSourceLocation(a.Source, param.Location().Row, param.Location().Column)
			dbg.ThrowError(
				fmt.Sprintf(
					"Function '%s' expects argument %d to be of type '%s', but got '%s'.",
					call.Name,
					i+1,
					function.Arguments[i].Type.Inspect(),
					paramType.Inspect(),
				),
				true,
			)
		}
	}

	return function.ReturnType
}
