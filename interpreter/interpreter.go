package interpreter

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
	"blom/interpreter/expressions"
	"fmt"
)

type Interpreter struct {
	source string
}

func New(source string) *Interpreter {
	return &Interpreter{
		source,
	}
}

func (intrepreter *Interpreter) Interpret(program *ast.Program) objects.Object {
	program.Body = append(program.Body, &ast.ReturnStatement{
		Value: &ast.FunctionCall{
			Name:       "main",
			Parameters: []ast.Expression{},
		},
	})

	return intrepreter.InterpretBlock(&ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}, env.New())
}

func (interpreter *Interpreter) InterpretBlock(body *ast.BlockStatement, environment *env.Environment) objects.Object {
	//fmt.Printf("Interpreting block %v\n", body.Location())

	envi := env.New(*environment)

	for _, stmt := range body.Body {
		value := interpreter.InterpretStatement(stmt, envi)

		//fmt.Printf("Interpreting statement %T, value: %v\n", stmt, value)

		switch stmt.(type) {
		case *ast.ReturnStatement:
			return value
		case *ast.IfStatement:
			if value != nil {
				return value
			}
		}
	}

	return nil
}

func (intrepreter *Interpreter) InterpretStatement(stmt ast.Statement, environment *env.Environment) objects.Object {
	switch stmt := stmt.(type) {
	case *ast.CharLiteralStatement:
		return &objects.CharacterObject{Value: stmt.Value}
	case *ast.StringLiteralStatement:
		return &objects.StringObject{Value: stmt.Value}
	case *ast.IntLiteralStatement:
		return &objects.IntObject{Value: stmt.Value}
	case *ast.FloatLiteralStatement:
		return &objects.FloatObject{Value: stmt.Value}
	case *ast.IdentifierLiteralStatement:
		return environment.FindVariable(stmt.Value)
	case *ast.BlockStatement:
		return intrepreter.InterpretBlock(stmt, environment)
	case *ast.FunctionDeclaration:
		expressions.InterpretFunctionDeclaration(intrepreter, environment, stmt)
	case *ast.BinaryExpression:
		return expressions.InterpretBinaryExpression(intrepreter, environment, stmt)
	case *ast.UnaryExpression:
		return expressions.InterpretUnaryExpression(intrepreter, environment, stmt)
	case *ast.DeclarationStatement:
		expressions.InterpretDeclarationStatement(intrepreter, environment, stmt)
	case *ast.ReturnStatement:
		return expressions.InterpretReturnStatement(intrepreter, environment, stmt)
	case *ast.IfStatement:
		return expressions.InterpretIfStatement(intrepreter, environment, stmt)
	case *ast.ForLoopStatement:
		expressions.InterpretForLoopStatement(intrepreter, environment, stmt)
	case *ast.WhileLoopStatement:
		expressions.InterpretWhileLoopStatement(intrepreter, environment, stmt)
	case *ast.FunctionCall:
		return expressions.InterpretFunctionCall(intrepreter, environment, stmt)
	case *ast.CompileTimeFunctionCall:
		return expressions.InterpretCompileTimeFunctionCall(intrepreter, environment, stmt)
	default:
		fmt.Printf("Unknown statement type: %T\n", stmt)
	}

	return nil
}

func (interpreter *Interpreter) Source() string {
	return interpreter.source
}
