package javascript

import (
	"blom/ast"
	"fmt"
)

func (t JavascriptTranspiler) TranspileReturnStatement(stmt *ast.Return) string {
	return fmt.Sprintf("return %s;", t.TranspileAndFunctionifyStatement(stmt.Value))
}
