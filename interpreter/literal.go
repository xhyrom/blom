package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
	"fmt"
)

func (t *Interpreter) interpretLiteral(literal ast.Statement) objects.Object {
	switch literal := literal.(type) {
	case *ast.IntLiteral:
		return &objects.IntObject{Value: int32(literal.Value)}
	case *ast.FloatLiteral:
		return &objects.FloatObject{Value: float32(literal.Value)}
	case *ast.CharLiteral:
		return &objects.CharacterObject{Value: literal.Value}
	case *ast.StringLiteral:
		return &objects.StringObject{Value: literal.Value}
	default:
		panic(fmt.Sprintf("'%T' is not a valid literal", literal))
	}
}
