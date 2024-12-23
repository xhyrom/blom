package javascript

import (
	"blom/ast"
	"blom/transpiler"
	"fmt"
	"strconv"
)

type JavascriptTranspiler struct{}

func (t JavascriptTranspiler) Transpile(program *ast.Program) (string, error) {
	program.Body = append(program.Body, &ast.FunctionCall{
		Name: "process.exit",
		Parameters: []ast.Expression{
			ast.FunctionCall{
				Name:       "main",
				Parameters: []ast.Expression{},
			},
		},
	})

	return t.TranspileBlock(ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}), nil
}

func (t JavascriptTranspiler) TranspileBlock(block ast.BlockStatement) string {
	result := ""

	for _, stmt := range block.Body {
		result += t.TranspileStatement(stmt)
	}

	return result
}

func (t JavascriptTranspiler) TranspileStatement(stmt ast.Statement) string {
	switch stmt := stmt.(type) {
	case *ast.IntLiteralStatement:
		return strconv.FormatInt(stmt.Value, 10)
	case ast.IntLiteralStatement:
		return strconv.FormatInt(stmt.Value, 10)
	case *ast.FloatLiteralStatement:
		return strconv.FormatFloat(stmt.Value, 'f', -1, 64)
	case ast.FloatLiteralStatement:
		return strconv.FormatFloat(stmt.Value, 'f', -1, 64)
	case *ast.IdentifierLiteralStatement:
		return stmt.Value
	case ast.IdentifierLiteralStatement:
		return stmt.Value
	case *ast.BlockStatement:
		return t.TranspileBlock(*stmt)
	case ast.BlockStatement:
		return t.TranspileBlock(stmt)
	case *ast.FunctionDeclaration:
		return t.TranspileFunctionDeclaration(stmt)
	case ast.FunctionDeclaration:
		return t.TranspileFunctionDeclaration(&stmt)
	case *ast.BinaryExpression:
		return t.TranspileBinaryExpression(stmt)
	case ast.BinaryExpression:
		return t.TranspileBinaryExpression(&stmt)
	case *ast.UnaryExpression:
		return t.TranspileUnaryExpression(stmt)
	case ast.UnaryExpression:
		return t.TranspileUnaryExpression(&stmt)
	case *ast.DeclarationStatement:
		return t.TranspileDeclarationStatement(stmt)
	case ast.DeclarationStatement:
		return t.TranspileDeclarationStatement(&stmt)
	case *ast.ReturnStatement:
		return t.TranspileReturnStatement(stmt)
	case ast.ReturnStatement:
		return t.TranspileReturnStatement(&stmt)
	case *ast.IfStatement:
		return t.TranspileIfStatement(stmt)
	case ast.IfStatement:
		return t.TranspileIfStatement(&stmt)
	case *ast.FunctionCall:
		return t.TranspileFunctionCall(stmt)
	case ast.FunctionCall:
		return t.TranspileFunctionCall(&stmt)
	default:
		fmt.Printf("Unknown statement type: %T\n", stmt)
	}

	return ""
}

var _ transpiler.Transpiler = (*JavascriptTranspiler)(nil)
