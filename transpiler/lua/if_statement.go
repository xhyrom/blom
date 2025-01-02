package lua

import (
	"blom/ast"
	"fmt"
)

func (t LuaTranspiler) TranspileIfStatement(statement *ast.If) string {
	result := fmt.Sprintf("if %s then\n", t.TranspileStatement(statement.Condition))

	result += t.TranspileStatement(statement.Then)

	if statement.Else != nil {
		result += "else\n"
		result += t.TranspileStatement(statement.Else)
	}

	result += "end\n"

	return result
}
