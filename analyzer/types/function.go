package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeFunctionDeclaration(function *ast.FunctionDeclaration) {
	for _, arg := range function.Arguments {
		a.Environment.Set(arg.Name, &Variable{Type: arg.Type})
	}

	for _, statement := range function.Body.Body {
		if statement.Kind() == ast.ReturnNode {
			ret := statement.(*ast.ReturnStatement)
			returnType := a.analyzeExpression(ret.Value)

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
			a.analyzeStatement(statement)
		}
	}

	a.Environment.SetFunction(function.Name, function)
}

func (a *TypeAnalyzer) analyzeFunctionCall(call *ast.FunctionCall) compiler.Type {
	function := a.Environment.GetFunction(call.Name)
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
		paramType := a.analyzeExpression(param)

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
