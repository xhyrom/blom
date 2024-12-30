package qbe

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"fmt"
)

func (c *Compiler) CompileFunctionCall(stmt *ast.FunctionCall, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	function := c.globalScope.GetFunction(stmt.Name)
	if function == nil {
		dbg := debug.NewSourceLocation(c.source, stmt.Loc.Row, stmt.Loc.Column)
		dbg.ThrowError(fmt.Sprintf("Function '%s' is not defined", stmt.Name), true)
	}

	name := fmt.Sprintf("%%tmp.%d", c.tempCounter)

	result := make([]string, 0)
	callResult := fmt.Sprintf("%s =%s call $%s(", name, c.StoreType(function.ReturnType), stmt.Name)

	for i := range function.Arguments {
		if i > 0 {
			callResult += ", "
		}

		if len(stmt.Parameters) <= i {
			dbg := debug.NewSourceLocation(c.source, stmt.Loc.Row, stmt.Loc.Column)
			dbg.ThrowError(fmt.Sprintf("Function '%s' expects %d arguments, but got %d", stmt.Name, len(function.Arguments), len(stmt.Parameters)), true)
		}

		param := stmt.Parameters[i]

		stat, identifier := c.CompileStatement(param, scope)

		for _, s := range stat {
			result = append(result, s)
		}

		callResult += identifier.String()
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

			stat, identifier := c.CompileStatement(param, scope)

			for _, s := range stat {
				result = append(result, s)
			}

			callResult += identifier.String()
		}
	}

	callResult += ")"

	result = append(result, callResult)

	result = append(result, "# ^ function call\n")

	return result, &QbeIdentifier{
		Name: name,
		Type: function.ReturnType,
	}
}
