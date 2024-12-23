package javascript

import (
	"blom/ast"
	"fmt"
)

func (t JavascriptTranspiler) TranspileDeclarationStatement(declaration *ast.DeclarationStatement) string {
	result := ""

	if declaration.Value.Kind() == ast.IfNode || declaration.Value.Kind() == ast.BlockNode {
		result += "(function() {\n"
		result += t.TranspileStatement(declaration.Value)
		result += "})()"
	} else {
		result += t.TranspileStatement(declaration.Value)
	}

	if declaration.Redeclaration {
		return fmt.Sprintf("%s = %s\n", declaration.Name, result)
	}

	return fmt.Sprintf("let %s = %s\n", declaration.Name, result)
}
