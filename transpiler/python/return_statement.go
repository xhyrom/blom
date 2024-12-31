package python

import (
	"blom/ast"
	"blom/env"
)

func (t PythonTranspiler) TranspileReturnStatement(statement *ast.ReturnStatement, environment *env.Scope, indent int) string {
	return "return " + t.TranspileStatement(statement.Value, environment, indent)
}
