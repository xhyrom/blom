package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
	"blom/tokens"
	"fmt"
)

func (t *Interpreter) interpretVariableDeclaration(statement *ast.VariableDeclarationStatement, function *ast.FunctionDeclaration, isReturn bool) {
	ty := statement.Type

	value := t.interpretStatement(statement.Value, function, &ty, isReturn)

	if ty != value.Type() {
		cnv := t.convertToType(value.Type(), ty, value)
		value = cnv
	}

	t.Scopes.Set(statement.Name, value)
}

func (t *Interpreter) interpretAssignment(statement *ast.Assignment, function *ast.FunctionDeclaration, isReturn bool) objects.Object {
	target := evaluateLeftSide(t, statement.Left, function)

	ty := target.Type()
	value := t.interpretStatement(statement.Right, function, &ty, isReturn)

	if ty != value.Type() {
		value = t.convertToType(value.Type(), ty, value)
	}

	target.SetValue(value.Value())

	return value
}

func evaluateLeftSide(t *Interpreter, left ast.Expression, function *ast.FunctionDeclaration) objects.Object {
	switch expr := left.(type) {
	case *ast.IdentifierLiteral:
		variable, exists := t.Scopes.GetValue(expr.Value)
		if !exists {
			panic("missing variable")
		}
		return variable

	case *ast.UnaryExpression:
		if expr.Operator != tokens.Asterisk {
			panic("unsupported unary operator")
		}

		operand := t.interpretStatement(expr.Operand, function, nil, false)
		return operand

	default:
		panic(fmt.Sprintf("unsupported left expression: %T", left))
	}
}
