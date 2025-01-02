package lua

import (
	"blom/ast"
	"blom/transpiler"
	"fmt"
	"strconv"
)

type LuaTranspiler struct{}

func (t LuaTranspiler) Transpile(program *ast.Program) (string, error) {
	program.Body = append(program.Body, &ast.FunctionCall{
		Name: "os.exit",
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

func (t LuaTranspiler) TranspileBlock(block ast.BlockStatement) string {
	result := "do\n"

	for _, stmt := range block.Body {
		result += t.TranspileStatement(stmt)
	}

	result += "\nend\n"

	return result
}

func (t LuaTranspiler) TranspileStatement(stmt ast.Statement) string {
	switch stmt := stmt.(type) {
	case *ast.CharLiteral:
		return "'" + string(stmt.Value) + "'"
	case ast.CharLiteral:
		return "'" + string(stmt.Value) + "'"
	case *ast.StringLiteral:
		return "\"" + stmt.Value + "\""
	case ast.StringLiteral:
		return "\"" + stmt.Value + "\""
	case *ast.IntLiteral:
		return strconv.FormatInt(stmt.Value, 10)
	case ast.IntLiteral:
		return strconv.FormatInt(stmt.Value, 10)
	case *ast.FloatLiteral:
		return strconv.FormatFloat(stmt.Value, 'f', -1, 64)
	case ast.FloatLiteral:
		return strconv.FormatFloat(stmt.Value, 'f', -1, 64)
	case *ast.IdentifierLiteral:
		return stmt.Value
	case ast.IdentifierLiteral:
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
	case *ast.VariableDeclarationStatement:
		return t.TranspileDeclarationStatement(stmt)
	case ast.VariableDeclarationStatement:
		return t.TranspileDeclarationStatement(&stmt)
	case *ast.ReturnStatement:
		return t.TranspileReturnStatement(stmt)
	case ast.ReturnStatement:
		return t.TranspileReturnStatement(&stmt)
	case *ast.If:
		return t.TranspileIfStatement(stmt)
	case ast.If:
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

var _ transpiler.Transpiler = (*LuaTranspiler)(nil)
