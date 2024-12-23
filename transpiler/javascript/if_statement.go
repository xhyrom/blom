package javascript

import (
	"blom/ast"
	"fmt"
)

func (t JavascriptTranspiler) TranspileIfStatement(statement *ast.IfStatement) string {
	result := fmt.Sprintf("if (%s) {\n", t.TranspileStatement(statement.Condition))

	result += t.TranspileStatement(statement.Then)

	if statement.Else != nil {
		result += "} else {\n"
		result += t.TranspileStatement(statement.Else)
	}

	result += "}\n"

	return result
}
