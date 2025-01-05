package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
	"fmt"
)

func (t *Interpreter) interpretLiteral(literal ast.Statement, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	switch literal := literal.(type) {
	case *ast.IdentifierLiteral:
		return interpretIdentifierLiteral(t, literal)
	case *ast.IntLiteral:
		return interpretIntLiteral(literal, function, vtype, isReturn)
	case *ast.FloatLiteral:
		return interpretFloatLiteral(literal, function, vtype, isReturn)
	case *ast.CharLiteral:
		return objects.NewCharacterObject(rune(literal.Value))
	case *ast.StringLiteral:
		return objects.NewStringObject(literal.Value)
	case *ast.BooleanLiteral:
		return objects.NewBooleanObject(literal.Value)
	}

	panic(fmt.Sprintf("'%T' is not a valid literal", literal))
}

func interpretIdentifierLiteral(t *Interpreter, literal *ast.IdentifierLiteral) objects.Object {
	variable, exists := t.Scopes.GetValue(literal.Value)
	if !exists {
		panic("missing variable")
	}

	return variable
}

func interpretIntLiteral(literal *ast.IntLiteral, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	var t ast.Type = ast.Int32
	if vtype != nil {
		t = *vtype
	}

	// casting (int to float)
	if t.IsFloatingPoint() {
		return interpretFloatLiteral(&ast.FloatLiteral{
			Value: float64(literal.Value),
			Loc:   literal.Loc,
		}, function, vtype, isReturn)
	}

	if isReturn {
		t = function.ReturnType
	}

	obj := objects.FromType(t)
	obj.SetValue(literal.Value)

	return obj
}

func interpretFloatLiteral(literal *ast.FloatLiteral, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	var t ast.Type = ast.Float32
	if vtype != nil {
		t = *vtype
	}

	// casting (float to int)
	if t.IsInteger() {
		return interpretIntLiteral(&ast.IntLiteral{
			Value: int64(literal.Value),
			Loc:   literal.Loc,
		}, function, vtype, isReturn)
	}

	if isReturn {
		t = function.ReturnType
	}

	obj := objects.FromType(t)
	obj.SetValue(literal.Value)

	return obj
}
