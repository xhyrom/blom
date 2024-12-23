package lua

import (
	"blom/ast"
	"fmt"
)

func (t LuaTranspiler) TranspileDeclarationStatement(declaration *ast.DeclarationStatement) string {
	if declaration.Redeclaration {
		return fmt.Sprintf("%s = %s\n", declaration.Name, t.TranspileStatement(declaration.Value))
	}

	return fmt.Sprintf("local %s = %s\n", declaration.Name, t.TranspileStatement(declaration.Value))
}
