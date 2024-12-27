package qbe

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (c *Compiler) CompileFunctionCall(stmt *ast.FunctionCall, env *env.Environment, ident int) string {
	function := env.GetFunction(stmt.Name)

	if function == nil {
		dbg := debug.NewSourceLocation(c.Source, stmt.Loc.Row, stmt.Loc.Column)
		dbg.ThrowError(fmt.Sprintf("Function '%s' is not defined", stmt.Name), true)
	}

	result := "call $" + stmt.Name + "("

	for i := range function.Arguments {
		if i > 0 {
			result += ", "
		}

		param := stmt.Parameters[i]
		result += "l " + c.CompileStatement(param, env, ident)
	}

	if function.Variadic {
		if len(function.Arguments) > 0 {
			result += ", "
		}

		result += "..."

		for i := len(function.Arguments); i < len(stmt.Parameters); i++ {
			if i > 0 {
				result += ", "
			}

			param := stmt.Parameters[i]
			result += "l " + c.CompileStatement(param, env, ident)
		}
	}

	result += ")"

	return result
}
