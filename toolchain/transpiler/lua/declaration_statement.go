package lua

import (
	"blom/ast"
	"fmt"
)

func (t LuaTranspiler) TranspileDeclarationStatement(declaration *ast.VariableDeclarationStatement) string {
	result := t.TranspileAndFunctionifyStatement(declaration.Value)

	if declaration.Redeclaration {
		return fmt.Sprintf("%s = %s\n", declaration.Name, result)
	}

	return fmt.Sprintf("local %s = %s\n", declaration.Name, result)
}
