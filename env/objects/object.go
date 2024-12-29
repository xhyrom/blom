package objects

import "blom/ast"

type Object interface {
	Type() ast.Type
	Inspect() string
	Add(other Object) Object
	Subtract(other Object) Object
	Multiply(other Object) Object
	Divide(other Object) Object
	Modulo(other Object) Object
	BitwiseAnd(other Object) Object
	BitwiseOr(other Object) Object
	BitwiseXor(other Object) Object
	LeftShift(other Object) Object
	RightShift(other Object) Object
	Equals(other Object) Object
	LessThan(other Object) Object
	LessThanOrEqual(other Object) Object
	GreaterThan(other Object) Object
	GreaterThanOrEqual(other Object) Object
}
