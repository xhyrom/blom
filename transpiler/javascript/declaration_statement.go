package javascript

import (
	"blom/ast"
	"fmt"
)

func (t JavascriptTranspiler) TranspileDeclarationStatement(declaration *ast.DeclarationStatement) string {
	if declaration.Redeclaration {
		return fmt.Sprintf("%s = %s\n", declaration.Name, t.TranspileStatement(declaration.Value))
	}

	return fmt.Sprintf("let %s = %s\n", declaration.Name, t.TranspileStatement(declaration.Value))
}
