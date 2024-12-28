package qbe

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (c *Compiler) CompileFunctionCall(stmt *ast.FunctionCall, ident int) ([]string, string) {
	function := c.Environment.GetFunction(stmt.Name)

	if function == nil {
		dbg := debug.NewSourceLocation(c.Source, stmt.Loc.Row, stmt.Loc.Column)
		dbg.ThrowError(fmt.Sprintf("Function '%s' is not defined", stmt.Name), true)
	}

	name := fmt.Sprintf("%%tmp.%d", c.Environment.TempCounter)
	c.Environment.TempCounter += 1

	result := make([]string, 0)
	callResult := fmt.Sprintf("%s =%s call $%s(", name, c.StoreType(function.ReturnType), stmt.Name)

	for i := range function.Arguments {
		if i > 0 {
			callResult += ", "
		}

		param := stmt.Parameters[i]

		stat, identifier := c.CompileStatement(param, ident)

		for _, s := range stat {
			result = append(result, s)
		}

		callResult += identifier
	}

	if function.Variadic {
		if len(function.Arguments) > 0 {
			callResult += ", "
		}

		callResult += "..."

		for i := len(function.Arguments); i < len(stmt.Parameters); i++ {
			if i > 0 {
				callResult += ", "
			}

			param := stmt.Parameters[i]

			stat, identifier := c.CompileStatement(param, ident)

			for _, s := range stat {
				result = append(result, s)
			}

			callResult += identifier
		}
	}

	callResult += ")"

	result = append(result, callResult)

	result = append(result, "# ^ function call\n")

	return result, name
}
