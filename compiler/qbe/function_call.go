package qbe

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (c *Compiler) CompileFunctionCall(stmt *ast.FunctionCall, ident int) string {
	function := c.Environment.GetFunction(stmt.Name)

	if function == nil {
		dbg := debug.NewSourceLocation(c.Source, stmt.Loc.Row, stmt.Loc.Column)
		dbg.ThrowError(fmt.Sprintf("Function '%s' is not defined", stmt.Name), true)
	}

	result := fmt.Sprintf("%%tmp.%d =%s call $%s(", c.Environment.TempCounter, c.StoreType(function.ReturnType), stmt.Name)

	for i := range function.Arguments {
		if i > 0 {
			result += ", "
		}

		param := stmt.Parameters[i]
		result += c.CompileStatement(param, ident)
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
			result += c.CompileStatement(param, ident)
		}
	}

	result += ")"

	return result
}
