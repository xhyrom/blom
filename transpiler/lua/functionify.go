package lua

import "blom/ast"

func (t LuaTranspiler) TranspileAndFunctionifyStatement(stmt ast.Statement) string {
	result := ""

	if stmt.Kind() == ast.IfNode || stmt.Kind() == ast.BlockNode {
		result += "(function()\n"
		result += t.TranspileStatement(stmt)
		result += "end)()"
	} else {
		result += t.TranspileStatement(stmt)
	}

	return result
}
