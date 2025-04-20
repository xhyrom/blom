package lua

import (
	"blom/ast"
	"fmt"
)

func (t LuaTranspiler) TranspileReturnStatement(stmt *ast.Return) string {
	return fmt.Sprintf("return %s", t.TranspileAndFunctionifyStatement(stmt.Value))
}
