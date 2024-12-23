package interpreter

import (
	"blom/ast"
	"blom/env"
	"blom/interpreter/expressions"
	"fmt"
)

type Interpreter struct{}

func New() *Interpreter {
	return &Interpreter{}
}

func (intrepreter *Interpreter) Interpret(program *ast.Program) env.Object {
	program.Body = append(program.Body, &ast.FunctionCall{
		Name:       "main",
		Parameters: []ast.Expression{},
	})

	return intrepreter.InterpretBlock(ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}, env.New())
}

func (interpreter *Interpreter) InterpretBlock(body ast.BlockStatement, environment *env.Environment) env.Object {
	envi := env.New(*environment)

	for _, stmt := range body.Body {
		value := interpreter.InterpretStatement(stmt, envi)

		fmt.Printf("Interpreting statement %T, value: %v\n", stmt, value)

		switch stmt.(type) {
		case *ast.FunctionCall, ast.FunctionCall:
			return value
		case *ast.ReturnStatement, ast.ReturnStatement:
			return value
		}
	}

	return nil
}

func (intrepreter *Interpreter) InterpretStatement(stmt ast.Statement, environment *env.Environment) env.Object {
	switch stmt := stmt.(type) {
	case *ast.IntLiteralStatement:
		return &env.IntegerObject{Value: stmt.Value}
	case ast.IntLiteralStatement:
		return &env.IntegerObject{Value: stmt.Value}
	case *ast.FloatLiteralStatement:
		return &env.FloatObject{Value: stmt.Value}
	case ast.FloatLiteralStatement:
		return &env.FloatObject{Value: stmt.Value}
	case *ast.IdentifierLiteralStatement:
		return environment.Get(stmt.Value)
	case ast.IdentifierLiteralStatement:
		return environment.Get(stmt.Value)
	case *ast.BlockStatement:
		return intrepreter.InterpretBlock(*stmt, environment)
	case ast.BlockStatement:
		return intrepreter.InterpretBlock(stmt, environment)
	case *ast.FunctionDeclaration:
		expressions.InterpretFunctionDeclaration(intrepreter, environment, stmt)
	case ast.FunctionDeclaration:
		expressions.InterpretFunctionDeclaration(intrepreter, environment, &stmt)
	case *ast.BinaryExpression:
		return expressions.InterpretBinaryExpression(intrepreter, environment, stmt)
	case ast.BinaryExpression:
		return expressions.InterpretBinaryExpression(intrepreter, environment, &stmt)
	case *ast.UnaryExpression:
		return expressions.InterpretUnaryExpression(intrepreter, environment, stmt)
	case ast.UnaryExpression:
		return expressions.InterpretUnaryExpression(intrepreter, environment, &stmt)
	case *ast.DeclarationStatement:
		expressions.InterpretDeclarationStatement(intrepreter, environment, stmt)
	case ast.DeclarationStatement:
		expressions.InterpretDeclarationStatement(intrepreter, environment, &stmt)
	case *ast.ReturnStatement:
		return expressions.InterpretReturnStatement(intrepreter, environment, stmt)
	case ast.ReturnStatement:
		return expressions.InterpretReturnStatement(intrepreter, environment, &stmt)
	case *ast.FunctionCall:
		return expressions.InterpretFunctionCall(intrepreter, environment, stmt)
	case ast.FunctionCall:
		return expressions.InterpretFunctionCall(intrepreter, environment, &stmt)
	default:
		fmt.Printf("Unknown statement type: %T\n", stmt)
	}

	return nil
}
