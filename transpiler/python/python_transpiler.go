package python

import (
	"blom/ast"
	"blom/env"
	"fmt"
	"strconv"
	"strings"
)

type PythonTranspiler struct{}

func (t PythonTranspiler) Transpile(program *ast.Program) (string, error) {
	program.Body = append(program.Body, &ast.FunctionCall{
		Name: "print",
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
	}, env.New(), 0), nil
}

func (t PythonTranspiler) TranspileBlock(block ast.BlockStatement, environment *env.Environment, indent int) string {
	indentation := strings.Repeat("    ", indent)

	env := env.New(*environment)

	result := ""
	if indent > 0 {
		result = "if True:\n"
	}

	for _, stmt := range block.Body {
		result += indentation + t.TranspileStatement(stmt, env, indent)
	}

	result += "\n"

	return result
}

func (t PythonTranspiler) TranspileStatement(stmt ast.Statement, environment *env.Environment, indent int) string {
	switch stmt := stmt.(type) {
	case *ast.CharLiteralStatement:
		return "'" + string(stmt.Value) + "'"
	case ast.CharLiteralStatement:
		return "'" + string(stmt.Value) + "'"
	case *ast.StringLiteralStatement:
		return "\"" + stmt.Value + "\""
	case ast.StringLiteralStatement:
		return "\"" + stmt.Value + "\""
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
		return t.TranspileBlock(*stmt, environment, indent+1)
	case ast.BlockStatement:
		return t.TranspileBlock(stmt, environment, indent+1)
	case *ast.FunctionDeclaration:
		return t.TranspileFunctionDeclaration(stmt, environment, indent)
	case ast.FunctionDeclaration:
		return t.TranspileFunctionDeclaration(&stmt, environment, indent)
	case *ast.BinaryExpression:
		return t.TranspileBinaryExpression(stmt, environment, indent)
	case ast.BinaryExpression:
		return t.TranspileBinaryExpression(&stmt, environment, indent)
	case *ast.UnaryExpression:
		return t.TranspileUnaryExpression(stmt, environment, indent)
	case ast.UnaryExpression:
		return t.TranspileUnaryExpression(&stmt, environment, indent)
	case *ast.VariableDeclarationStatement:
		return t.TranspileDeclarationStatement(stmt, environment, indent)
	case ast.VariableDeclarationStatement:
		return t.TranspileDeclarationStatement(&stmt, environment, indent)
	case *ast.ReturnStatement:
		return t.TranspileReturnStatement(stmt, environment, indent)
	case ast.ReturnStatement:
		return t.TranspileReturnStatement(&stmt, environment, indent)
	case *ast.IfStatement:
		return t.TranspileIfStatement(stmt, environment, indent)
	case ast.IfStatement:
		return t.TranspileIfStatement(&stmt, environment, indent)
	case *ast.FunctionCall:
		return t.TranspileFunctionCall(stmt, environment, indent)
	case ast.FunctionCall:
		return t.TranspileFunctionCall(&stmt, environment, indent)
	default:
		fmt.Printf("Unknown statement type: %T\n", stmt)
	}

	return ""
}

// TODO
//var _ transpiler.Transpiler = (*PythonTranspiler)(nil)
