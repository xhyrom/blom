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
	Context         *AnalysisContext
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
		Context:         NewAnalysisContext(),
	}
}

func (a *TypeAnalyzer) Analyze() {
	for _, node := range a.Program.Body {
		a.analyzeStatement(node)
	}
}

func (a *TypeAnalyzer) analyzeStatement(statement ast.Node) (ast.Type, bool) {
	switch statement := statement.(type) {
	case *ast.FunctionDeclaration:
		a.analyzeFunctionDeclaration(statement)
	case *ast.VariableDeclaration:
		a.analyzeVariableDeclarationStatement(statement)
	case *ast.WhileLoop:
		a.analyzeWhileLoopStatement(statement)
	case *ast.FunctionCall, *ast.MethodCall, *ast.InfixCall:
		a.analyzeCall(statement.(ast.Call))
	default:
		if statement.Kind() != ast.IfNode && statement.Kind() != ast.AssignmentNode && statement.Kind() != ast.BlockNode {
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

func (a *TypeAnalyzer) analyzeExpression(expression ast.Node) ast.Type {
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
	case *ast.If:
		ifExpression := expression.(*ast.If)
		return a.analyzeIf(ifExpression)
	case *ast.Assignment:
		assignmentExpression := expression.(*ast.Assignment)
		return a.analyzeAssignment(assignmentExpression)
	case *ast.FunctionCall, *ast.MethodCall, *ast.InfixCall:
		return a.analyzeCall(expression.(ast.Call))
	case *ast.Block:
		blockStatement := expression.(*ast.Block)
		return a.analyzeBlock(blockStatement)
	}

	return ast.Void
}

func (a *TypeAnalyzer) canBeImplicitlyCast(from ast.Type, to ast.Type) bool {
	if from.IsFunction() && to.IsPointer() && to.(ast.PointerType).Dereference() == ast.Void {
		return true
	}

	if from.IsPointer() && from.(ast.PointerType).Dereference() == ast.Void && to.IsFunction() {
		return true
	}

	if from.IsFunction() && to.IsFunction() {
		var fromFunction ast.FunctionType
		var toFunction ast.FunctionType

		if from.IsPointer() {
			from = from.(ast.PointerType).Dereference()
			fromFunction = from.(ast.FunctionType)
		} else {
			fromFunction = from.(ast.FunctionType)
		}

		if to.IsPointer() {
			to = to.(ast.PointerType).Dereference()
			toFunction = to.(ast.FunctionType)
		} else {
			toFunction = to.(ast.FunctionType)
		}

		if fromFunction.ReturnType != toFunction.ReturnType {
			return false
		}

		if len(fromFunction.Arguments) != len(toFunction.Arguments) {
			return false
		}

		for i, fromArg := range fromFunction.Arguments {
			toArg := toFunction.Arguments[i]

			if fromArg != toArg {
				return false
			}
		}

		return true
	}

	if (from.IsFunction() && !to.IsFunction()) || (!from.IsFunction() && to.IsFunction()) {
		return false
	}

	if from.IsPointer() && from.(ast.PointerType).Dereference() == ast.Void {
		return true
	}

	if to.IsPointer() && to.(ast.PointerType).Dereference() == ast.Void {
		return true
	}

	if from == to {
		return true
	}

	if from.IsNumeric() && to.IsNumeric() {
		fromWeight := from.Weight()
		toWeight := to.Weight()

		return fromWeight <= toWeight && fromWeight <= uint8(ast.Float64) && to.Weight() <= uint8(ast.Float64)
	}

	return false
}

func (a *TypeAnalyzer) GetContext() *AnalysisContext {
	return a.Context
}
