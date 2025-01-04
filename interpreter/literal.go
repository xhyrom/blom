package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
	"fmt"
)

func (t *Interpreter) interpretLiteral(literal ast.Statement) objects.Object {
	switch literal := literal.(type) {
	case *ast.IdentifierLiteral:
		return interpretIdentifierLiteral(t, literal)
	case *ast.IntLiteral:
		return objects.NewIntObject(int32(literal.Value))
	case *ast.FloatLiteral:
		return objects.NewFloatObject(float32(literal.Value))
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
