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
	Environment *env.Environment[*Variable]
}

type Variable struct {
	Type compiler.Type
}

func New(file string, program *ast.Program) *TypeAnalyzer {
	return &TypeAnalyzer{
		Source:      file,
		Program:     program,
		Environment: env.New[*Variable](),
	}
}

func (a *TypeAnalyzer) Analyze() {
	for _, statement := range a.Program.Body {
		a.analyzeStatement(statement)
	}
}

func (a *TypeAnalyzer) analyzeStatement(statement ast.Statement) {
	switch statement := statement.(type) {
	case *ast.FunctionDeclaration:
		a.analyzeFunctionDeclaration(statement)
	case *ast.VariableDeclarationStatement:
		a.analyzeVariableDeclarationStatement(statement)
	case *ast.AssignmentStatement:
		a.analyzeAssignmentStatement(statement)
	case *ast.IfStatement: // if is statement but also an expression
		a.analyzeIfExpression(statement)
	case *ast.WhileLoopStatement:
		a.analyzeWhileLoopStatement(statement)
	case *ast.FunctionCall:
		a.analyzeFunctionCall(statement)
	case *ast.BlockStatement:
		for _, statement := range statement.Body {
			a.analyzeStatement(statement)
		}
	default:
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

		a.analyzeExpression(statement)
	}
}

func (a *TypeAnalyzer) analyzeExpression(expression ast.Expression) compiler.Type {
	switch expression.(type) {
	case *ast.IntLiteralStatement:
		return compiler.Word
	case *ast.FloatLiteralStatement:
		return compiler.Double
	case *ast.StringLiteralStatement:
		return compiler.String
	case *ast.CharLiteralStatement:
		return compiler.Char
	case *ast.IdentifierLiteralStatement:
		identifier := expression.(*ast.IdentifierLiteralStatement)
		return a.analyzeIdentifier(identifier)
		//case *ast.BooleanLiteralStatement:
		//	return compiler.Boolean
	case *ast.BinaryExpression:
		binaryExpression := expression.(*ast.BinaryExpression)
		return a.analyzeBinaryExpression(binaryExpression)
	case *ast.UnaryExpression:
		unaryExpression := expression.(*ast.UnaryExpression)
		return a.analyzeUnaryExpression(unaryExpression)
	case *ast.IfStatement: // if is statement but also an expression
		ifExpression := expression.(*ast.IfStatement)
		return a.analyzeIfExpression(ifExpression)
	case *ast.FunctionCall:
		functionCall := expression.(*ast.FunctionCall)
		return a.analyzeFunctionCall(functionCall)
	}

	return compiler.Void
}
