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
		intLiteral := expression.(*ast.IntLiteral)
		intLiteral.Type = ast.Int32

		return ast.Int32
	case *ast.FloatLiteral:
		floatLiteral := expression.(*ast.FloatLiteral)
		floatLiteral.Type = ast.Float32

		return ast.Float32
	case *ast.StringLiteral:
		stringLiteral := expression.(*ast.StringLiteral)
		stringLiteral.Type = ast.String

		return ast.String
	case *ast.CharLiteral:
		charLiteral := expression.(*ast.CharLiteral)
		charLiteral.Type = ast.Char

		return ast.Char
	case *ast.IdentifierLiteral:
		identifier := expression.(*ast.IdentifierLiteral)
		typ := a.analyzeIdentifier(identifier, scope)

		identifier.Type = typ

		return typ
	case *ast.BooleanLiteral:
		booleanLiteral := expression.(*ast.BooleanLiteral)
		booleanLiteral.Type = ast.Boolean

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

func (a *TypeAnalyzer) setExpressionType(expression ast.Expression, newType ast.Type) ast.Type {
	switch expression.(type) {
	case *ast.IntLiteral:
		intLiteral := expression.(*ast.IntLiteral)
		intLiteral.Type = newType

		return newType
	case *ast.FloatLiteral:
		floatLiteral := expression.(*ast.FloatLiteral)
		floatLiteral.Type = newType

		return newType
	case *ast.StringLiteral:
		stringLiteral := expression.(*ast.StringLiteral)
		stringLiteral.Type = newType

		return newType
	case *ast.CharLiteral:
		charLiteral := expression.(*ast.CharLiteral)
		charLiteral.Type = newType

		return newType
	case *ast.IdentifierLiteral:
		identifier := expression.(*ast.IdentifierLiteral)
		identifier.Type = newType

		return newType
	}

	panic("unreachable")
}
