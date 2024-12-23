package lua

import (
	"blom/ast"
	"fmt"
)

func (t LuaTranspiler) TranspileDeclarationStatement(declaration *ast.DeclarationStatement) string {
	result := ""

	if declaration.Value.Kind() == ast.IfNode || declaration.Value.Kind() == ast.BlockNode {
		result += "(function()\n"
		result += t.TranspileStatement(declaration.Value)
		result += "end)()"
	} else {
		result += t.TranspileStatement(declaration.Value)
	}

	if declaration.Redeclaration {
		return fmt.Sprintf("%s = %s\n", declaration.Name, result)
	}

	return fmt.Sprintf("local %s = %s\n", declaration.Name, result)
}
