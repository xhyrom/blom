package lua

import (
	"blom/ast"
	"fmt"
)

func (t LuaTranspiler) TranspileReturnStatement(stmt *ast.ReturnStatement) string {
	return fmt.Sprintf("return %s", t.TranspileStatement(stmt.Value))
}
