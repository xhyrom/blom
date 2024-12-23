package lua

import (
	"blom/ast"
	"fmt"
)

func (t LuaTranspiler) TranspileDeclarationStatement(declaration *ast.DeclarationStatement) string {
	return fmt.Sprintf("local %s = %s\n", declaration.Name, t.TranspileStatement(declaration.Value))
}
