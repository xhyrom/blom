package python

import (
	"blom/ast"
	"blom/env"
)

func (t PythonTranspiler) TranspileIfStatement(statement *ast.IfStatement, environment *env.Scope, indent int) string {
	result := "if " + t.TranspileStatement(statement.Condition, environment, indent) + ":\n"

	result += t.TranspileStatement(statement.Then, environment, indent)

	if statement.Else != nil {
		result += "else:\n"
		result += t.TranspileStatement(statement.Else, environment, indent)
	}

	result += "\n"

	return result
}
