package types

import (
	"blom/ast"
	"blom/debug"
	"blom/scope"
	"fmt"
)

type TypeAnalyzer struct {
	Source    string
	Program   *ast.Program
	Scopes    []scope.Scope[*Variable]
	Functions scope.Scope[*ast.FunctionDeclaration]
}

type Variable struct {
	Type ast.Type
}

func New(file string, program *ast.Program, functions map[string]*ast.FunctionDeclaration) *TypeAnalyzer {
	funcs := scope.New[*ast.FunctionDeclaration]()

	for _, function := range functions {
		funcs.Set(function.Name, function)
	}

	return &TypeAnalyzer{
		Source:    file,
		Program:   program,
		Scopes:    make([]scope.Scope[*Variable], 0),
		Functions: funcs,
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
	case *ast.AssignmentStatement:
		a.analyzeAssignmentStatement(statement)
	case *ast.WhileLoopStatement:
		a.analyzeWhileLoopStatement(statement)
	case *ast.FunctionCall:
		a.analyzeFunctionCall(statement)
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
	case *ast.IfStatement: // if is statement but also an expression
		ifExpression := expression.(*ast.IfStatement)
		return a.analyzeIfExpression(ifExpression)
	case *ast.FunctionCall:
		functionCall := expression.(*ast.FunctionCall)
		return a.analyzeFunctionCall(functionCall)
	case *ast.CompileTimeFunctionCall:
		compileTimeFunctionCall := expression.(*ast.CompileTimeFunctionCall)
		return a.analyzeCompileTimeFunctionCall(compileTimeFunctionCall)
	case *ast.BlockStatement:
		blockStatement := expression.(*ast.BlockStatement)
		return a.analyzeBlock(blockStatement)
	}

	return ast.Void
}

func (a *TypeAnalyzer) createVariable(name string, variable *Variable) {
	a.Scopes[len(a.Scopes)-1].Set(name, variable)
}

func (a *TypeAnalyzer) getVariable(name string) *Variable {
	for i := len(a.Scopes) - 1; i >= 0; i-- {
		value, exists := a.Scopes[i].Get(name)
		if exists {
			return value
		}
	}

	return nil
}
