package types

import (
	"blom/analyzer/manager"
	"blom/ast"
	"blom/debug"
	"blom/scope"
	"fmt"
)

type TypeAnalyzer struct {
	Source          string
	Program         *ast.Program
	Scopes          *scope.Scopes[*Variable]
	FunctionManager *manager.FunctionManager
}

type Variable struct {
	Type ast.Type
}

func New(file string, program *ast.Program, functionManager *manager.FunctionManager) *TypeAnalyzer {
	return &TypeAnalyzer{
		Source:          file,
		Program:         program,
		Scopes:          scope.NewScopes[*Variable](),
		FunctionManager: functionManager,
	}
}

func (a *TypeAnalyzer) Analyze() {
	for _, statement := range a.Program.Body {
		a.analyzeStatement(statement)
	}
}

func (a *TypeAnalyzer) analyzeStatement(statement ast.Statement) (ast.Type, bool) {
	switch statement := statement.(type) {
	case *ast.FunctionDeclaration:
		a.analyzeFunctionDeclaration(statement)
	case *ast.VariableDeclarationStatement:
		a.analyzeVariableDeclarationStatement(statement)
	case *ast.WhileLoopStatement:
		a.analyzeWhileLoopStatement(statement)
	case *ast.FunctionCall:
		a.analyzeFunctionCall(statement)
	default:
		if statement.Kind() != ast.IfNode && statement.Kind() != ast.AssignmentNode {
			dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column)
			dbg.ThrowWarning(
				fmt.Sprintf(
					"The statement '%T' has no effect on the program's behavior.",
					statement.Kind(),
				),
				true,
				debug.NewHint(
					"Consider removing this statement as it does not affect the program's behavior.",
					"",
				),
			)
		}

		return a.analyzeExpression(statement), true
	}

	return ast.Void, false
}

func (a *TypeAnalyzer) analyzeExpression(expression ast.Expression) ast.Type {
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
		return a.analyzeIdentifier(identifier)
	case *ast.BooleanLiteral:
		return ast.Boolean
	case *ast.BinaryExpression:
		binaryExpression := expression.(*ast.BinaryExpression)
		return a.analyzeBinaryExpression(binaryExpression)
	case *ast.UnaryExpression:
		unaryExpression := expression.(*ast.UnaryExpression)
		return a.analyzeUnaryExpression(unaryExpression)
	case *ast.If: // if is statement but also an expression
		ifExpression := expression.(*ast.If)
		return a.analyzeIf(ifExpression)
	case *ast.Assignment:
		assignmentExpression := expression.(*ast.Assignment)
		return a.analyzeAssignment(assignmentExpression)
	case *ast.FunctionCall:
		functionCall := expression.(*ast.FunctionCall)
		return a.analyzeFunctionCall(functionCall)
	case *ast.BuiltinFunctionCall:
		compileTimeFunctionCall := expression.(*ast.BuiltinFunctionCall)
		return a.analyzeBuiltinFunctionCall(compileTimeFunctionCall)
	case *ast.BlockStatement:
		blockStatement := expression.(*ast.BlockStatement)
		return a.analyzeBlock(blockStatement)
	}

	return ast.Void
}

func (a *TypeAnalyzer) canBeImplicitlyCast(from ast.Type, to ast.Type) bool {
	if from == to {
		return true
	}

	return from < to && from <= ast.Float64 && to <= ast.Float64
}
