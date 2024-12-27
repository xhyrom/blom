package qbe

import (
	"blom/ast"
	"strconv"
	"strings"
)

type Compiler struct{}

func New() *Compiler {
	return &Compiler{}
}

func (c Compiler) Compile(program *ast.Program) (string, error) {
	return c.CompileBlock(ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}, 0), nil
}

func (c Compiler) CompileBlock(block ast.BlockStatement, ident int) string {
	result := ""
	identation := strings.Repeat("    ", ident)

	for _, stmt := range block.Body {
		result += identation + c.CompileStatement(stmt, ident) + "\n"
	}

	return result
}

func (c Compiler) CompileStatement(stmt ast.Statement, ident int) string {
	switch stmt := stmt.(type) {
	case *ast.IntLiteralStatement:
		return strconv.FormatInt(int64(stmt.Value), 10)
	case *ast.DeclarationStatement:
		return c.CompileDeclarationStatement(stmt, ident+1)
	case *ast.FunctionCall:
		return c.CompileFunctionCall(stmt, ident+1)
	case *ast.FunctionDeclaration:
		return c.CompileFunctionDeclaration(stmt, ident+1)
	case *ast.ReturnStatement:
		return c.CompileReturnStatement(stmt, ident+1)
	case *ast.BlockStatement:
		return c.CompileBlock(*stmt, ident+1)
	case *ast.BinaryExpression:
		return c.CompileBinaryExpression(stmt, ident+1)
	}

	return ""
}
