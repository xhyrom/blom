package types

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/env"
	"fmt"
)

type TypeAnalyzer struct {
	Source      string
	Program     *ast.Program
	GlobalScope *env.Environment[*Variable]
}

type Variable struct {
	Type compiler.Type
}

func New(file string, program *ast.Program, functions map[string]*ast.FunctionDeclaration) *TypeAnalyzer {
	globalScope := env.New[*Variable]()

	for _, function := range functions {
		globalScope.SetFunction(function.Name, function)
	}

	return &TypeAnalyzer{
		Source:      file,
		Program:     program,
		GlobalScope: globalScope,
	}
}

func (a *TypeAnalyzer) Analyze() {
	for _, statement := range a.Program.Body {
		a.analyzeStatement(statement, a.GlobalScope)
	}
}

func (a *TypeAnalyzer) analyzeStatement(statement ast.Statement, scope *env.Environment[*Variable]) (compiler.Type, bool) {
	switch statement := statement.(type) {
	case *ast.FunctionDeclaration:
		a.analyzeFunctionDeclaration(statement)
	case *ast.VariableDeclarationStatement:
		a.analyzeVariableDeclarationStatement(statement, scope)
	case *ast.AssignmentStatement:
		a.analyzeAssignmentStatement(statement, scope)
	case *ast.WhileLoopStatement:
		a.analyzeWhileLoopStatement(statement, scope)
	case *ast.FunctionCall:
		a.analyzeFunctionCall(statement, scope)
	default:
		if statement.Kind() != ast.IfNode {
			dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column)
			dbg.ThrowWarning(
				fmt.Sprintf(
					"The statement '%s' has no effect on the program's behavior.",
					statement.Kind(),
				),
				true,
				debug.NewHint(
					"Consider removing this statement as it does not affect the program's behavior.",
					"",
				),
			)
		}

		return a.analyzeExpression(statement, scope), true
	}

	return compiler.Void, false
}

func (a *TypeAnalyzer) analyzeExpression(expression ast.Expression, scope *env.Environment[*Variable]) compiler.Type {
	switch expression.(type) {
	case *ast.IntLiteralStatement:
		intLiteral := expression.(*ast.IntLiteralStatement)
		intLiteral.Type = compiler.Word

		return compiler.Word
	case *ast.FloatLiteralStatement:
		floatLiteral := expression.(*ast.FloatLiteralStatement)
		floatLiteral.Type = compiler.Double

		return compiler.Double
	case *ast.StringLiteralStatement:
		stringLiteral := expression.(*ast.StringLiteralStatement)
		stringLiteral.Type = compiler.String

		return compiler.String
	case *ast.CharLiteralStatement:
		charLiteral := expression.(*ast.CharLiteralStatement)
		charLiteral.Type = compiler.Char

		return compiler.Char
	case *ast.IdentifierLiteralStatement:
		identifier := expression.(*ast.IdentifierLiteralStatement)
		typ := a.analyzeIdentifier(identifier, scope)

		identifier.Type = typ

		return typ
	case *ast.BooleanLiteralStatement:
		booleanLiteral := expression.(*ast.BooleanLiteralStatement)
		booleanLiteral.Type = compiler.Boolean

		return compiler.Boolean
	case *ast.BinaryExpression:
		binaryExpression := expression.(*ast.BinaryExpression)
		return a.analyzeBinaryExpression(binaryExpression, scope)
	case *ast.UnaryExpression:
		unaryExpression := expression.(*ast.UnaryExpression)
		return a.analyzeUnaryExpression(unaryExpression, scope)
	case *ast.IfStatement: // if is statement but also an expression
		ifExpression := expression.(*ast.IfStatement)
		return a.analyzeIfExpression(ifExpression, scope)
	case *ast.FunctionCall:
		functionCall := expression.(*ast.FunctionCall)
		return a.analyzeFunctionCall(functionCall, scope)
	case *ast.BlockStatement:
		blockStatement := expression.(*ast.BlockStatement)
		return a.analyzeBlock(blockStatement, scope)
	}

	return compiler.Void
}
