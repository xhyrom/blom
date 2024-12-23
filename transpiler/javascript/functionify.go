package javascript

import "blom/ast"

func (t JavascriptTranspiler) TranspileAndFunctionifyStatement(stmt ast.Statement) string {
	result := ""

	if stmt.Kind() == ast.IfNode || stmt.Kind() == ast.BlockNode {
		result += "(function() {\n"
		result += t.TranspileStatement(stmt)
		result += "})()"
	} else {
		result += t.TranspileStatement(stmt)
	}

	return result
}
