package types

import (
	"blom/ast"
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
	Type ast.Type
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

func (a *TypeAnalyzer) analyzeStatement(statement ast.Statement, scope *env.Environment[*Variable]) (ast.Type, bool) {
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

	return ast.Void, false
}

func (a *TypeAnalyzer) analyzeExpression(expression ast.Expression, scope *env.Environment[*Variable]) ast.Type {
	switch expression.(type) {
	case *ast.IntLiteral:
		return ast.Int32
	case *ast.FloatLiteral:
		return ast.Float32
	case *ast.StringLiteral:
		return ast.String
	case *ast.CharLiteral:
		return ast.Char
	case *ast.IdentifierLiteral:
		identifier := expression.(*ast.IdentifierLiteral)

		return a.analyzeIdentifier(identifier, scope)
	case *ast.BooleanLiteral:
		return ast.Boolean
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
	case *ast.CompileTimeFunctionCall:
		compileTimeFunctionCall := expression.(*ast.CompileTimeFunctionCall)
		return a.analyzeCompileTimeFunctionCall(compileTimeFunctionCall)
	case *ast.BlockStatement:
		blockStatement := expression.(*ast.BlockStatement)
		return a.analyzeBlock(blockStatement, scope)
	}

	return ast.Void
}
