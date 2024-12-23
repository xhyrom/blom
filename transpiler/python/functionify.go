package python

import (
	"blom/ast"
	"blom/env"
)

func (t PythonTranspiler) TranspileAndFunctionifyStatement(stmt ast.Statement, environment *env.Environment, indent int) string {
	return t.TranspileStatement(stmt, environment, indent)
}
