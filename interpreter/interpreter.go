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

func (intrepreter *Interpreter) Interpret(prg *ast.Program, argc int64) objects.Object {
	// copy program body to avoid modifying the original
	var program *ast.Program = &ast.Program{
		Body: make([]ast.Statement, 0),
	}

	for _, stmt := range prg.Body {
		program.Body = append(program.Body, stmt)
	}

	program.Body = append(program.Body, &ast.FunctionCall{
		Name:       "main",
		Parameters: []ast.Expression{},
	})

	return intrepreter.InterpretBlock(&ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}, env.New[objects.Object]())
}

func (interpreter *Interpreter) InterpretBlock(body *ast.BlockStatement, environment *env.Scope[objects.Object]) objects.Object {
	envi := env.New(*environment)

	for _, stmt := range body.Body {
		value := interpreter.InterpretStatement(stmt, envi)

		switch stmt.(type) {
		case *ast.ReturnStatement:
			return value
		case *ast.If:
			if value != nil {
				return value
			}
		}
	}

	return nil
}

func (intrepreter *Interpreter) InterpretStatement(stmt ast.Statement, environment *env.Scope[objects.Object]) objects.Object {
	switch stmt := stmt.(type) {
	case *ast.CharLiteral:
		return &objects.CharacterObject{Value: stmt.Value}
	case *ast.StringLiteral:
		return &objects.StringObject{Value: stmt.Value}
	case *ast.IntLiteral:
		return &objects.IntObject{Value: int32(stmt.Value)}
	case *ast.FloatLiteral:
		return &objects.FloatObject{Value: float32(stmt.Value)}
	case *ast.BooleanLiteral:
		var value int32 = 0
		if stmt.Value {
			value = 1
		}

		return &objects.IntObject{Value: value}
	case *ast.IdentifierLiteral:
		variable, found := environment.FindVariable(stmt.Value)
		if !found {
			panic(fmt.Sprintf("Variable %s not found", stmt.Value))
		}

		return variable
	case *ast.BlockStatement:
		return intrepreter.InterpretBlock(stmt, environment)
	case *ast.FunctionDeclaration:
		expressions.InterpretFunctionDeclaration(intrepreter, environment, stmt)
	case *ast.BinaryExpression:
		return expressions.InterpretBinaryExpression(intrepreter, environment, stmt)
	case *ast.UnaryExpression:
		return expressions.InterpretUnaryExpression(intrepreter, environment, stmt)
	case *ast.VariableDeclarationStatement:
		expressions.InterpretDeclarationStatement(intrepreter, environment, stmt)
	case *ast.Assignment:
		expressions.InterpretAssignmentStatement(intrepreter, environment, stmt)
	case *ast.ReturnStatement:
		return expressions.InterpretReturnStatement(intrepreter, environment, stmt)
	case *ast.If:
		return expressions.InterpretIfStatement(intrepreter, environment, stmt)
	case *ast.WhileLoopStatement:
		expressions.InterpretWhileLoopStatement(intrepreter, environment, stmt)
	case *ast.FunctionCall:
		return expressions.InterpretFunctionCall(intrepreter, environment, stmt)
	case *ast.BuiltinFunctionCall:
		return expressions.InterpretCompileTimeFunctionCall(intrepreter, environment, stmt)
	default:
		fmt.Printf("Unknown statement type: %T\n", stmt)
	}

	return nil
}

func (interpreter *Interpreter) Source() string {
	return interpreter.source
}
