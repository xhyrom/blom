package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeFunctionDeclaration(function *ast.FunctionDeclaration) {
	a.Scopes.Append()

	for _, param := range function.Params {
		a.Scopes.Set(param.Name.Value, &Variable{Type: param.Type})
	}

	if function.HasAnnotation(ast.Native) {
		return
	}

	for _, statement := range function.Block.Body {
		if statement.Kind() == ast.ReturnNode {
			ret := statement.(*ast.Return)
			returnType := a.analyzeExpression(ret.Value)

			if returnType != function.Return && (ret.Value.Kind() != ast.IntNode && !a.canBeImplicitlyCast(returnType, function.Return)) {
				dbg := debug.NewSourceLocationFromNode(a.Source, ret)
				dbg.ThrowError(
					fmt.Sprintf(
						"Function '%s' returns '%s', but declared to return '%s'",
						function.Name,
						returnType,
						function.Return,
					),
					true,
				)
			}
		} else {
			a.analyzeStatement(statement)
		}
	}

	fun, _ := a.FunctionManager.GetByDeclaration(function)
	suffix := a.FunctionManager.GetFuncNameSuffix(function)
	if suffix != "" {
		function.Name.Value = function.Name.Value + "." + suffix
		fun.Suffix = suffix
	}

	a.Scopes.Pop()
}
