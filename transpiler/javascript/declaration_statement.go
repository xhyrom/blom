package javascript

import (
	"blom/ast"
	"fmt"
)

func (t JavascriptTranspiler) TranspileDeclarationStatement(declaration *ast.DeclarationStatement) string {
	return fmt.Sprintf("local %s = %s\n", declaration.Name, t.TranspileStatement(declaration.Value))
}
