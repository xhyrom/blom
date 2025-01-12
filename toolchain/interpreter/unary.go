package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
	"blom/tokens"
	"fmt"
)

func (t *Interpreter) interpretUnaryExpression(expression *ast.UnaryExpression, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	switch expression.Operator {
	case tokens.Plus: // unary plus
		return t.interpretStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: 1,
			},
			Loc:         expression.Operand.Location(),
			Operator:    tokens.Asterisk, // multiply by 1
			OperatorLoc: expression.Operand.Location(),
		}, function, vtype, isReturn)
	case tokens.Minus: // unary minus
		return t.interpretStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: -1,
			},
			Loc:         expression.Operand.Location(),
			Operator:    tokens.Asterisk, // multiply by -1
			OperatorLoc: expression.Operand.Location(),
		}, function, vtype, isReturn)
	case tokens.Tilde: // bitwise not
		return t.interpretStatement(&ast.BinaryExpression{
			Left: expression.Operand,
			Right: &ast.IntLiteral{
				Value: -1,
			},
			Loc:         expression.Operand.Location(),
			Operator:    tokens.CircumflexAccent, // bitwise xor
			OperatorLoc: expression.Operand.Location(),
		}, function, vtype, isReturn)
	case tokens.Ampersand: // address of
		return interpretAddressOf(t, expression.Operand, function, vtype)
	case tokens.Asterisk: // dereference
		return interpretDereference(t, expression.Operand, function, vtype)
	}

	panic(fmt.Sprintf("unknown unary operator: %s", expression.Operator))
}

func interpretAddressOf(t *Interpreter, expression ast.Expression, function *ast.FunctionDeclaration, vtype *ast.Type) objects.Object {
	if ident, ok := expression.(*ast.IdentifierLiteral); ok {
		variable, exists := t.Scopes.GetValue(ident.Value)
		if !exists {
			panic(fmt.Sprintf("undefined variable: %s", ident.Value))
		}

		return objects.NewPointerObject(variable)
	}

	panic(fmt.Sprintf("cannot take address of %T", expression))
}

func interpretDereference(t *Interpreter, expression ast.Expression, function *ast.FunctionDeclaration, vtype *ast.Type) objects.Object {
	ptr := t.interpretStatement(expression, function, vtype, false)

	if pointerObj, ok := ptr.(*objects.PointerObject); ok {
		return pointerObj.Value().(objects.Object)
	}

	panic(fmt.Sprintf("cannot dereference non-pointer type: %T", ptr))
}
