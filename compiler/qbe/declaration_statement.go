package qbe

import "blom/ast"

func (c Compiler) CompileDeclarationStatement(stmt *ast.DeclarationStatement, ident int) string {
	result := "%" + stmt.Name

	if stmt.Value != nil {
		result += " =" + stmt.Type.String() + " " + c.CompileStatement(stmt.Value, ident)
	}

	return result
}
