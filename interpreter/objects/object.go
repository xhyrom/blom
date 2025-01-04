package objects

import "blom/ast"

type Object interface {
	Type() ast.Type
	Inspect() string
	Value() interface{}
	SetValue(value interface{})
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

func FromType(t ast.Type) Object {
	switch t {
	case ast.Int8:
		return &ByteObject{}
	case ast.UnsignedInt8:
		return &UnsignedByteObject{}
	case ast.Int16:
		return &ShortObject{}
	case ast.UnsignedInt16:
	case ast.Int32:
		return &IntObject{}
	case ast.UnsignedInt32:
		return &UnsignedIntObject{}
	case ast.Int64:
		return &LongObject{}
	case ast.UnsignedInt64:
		return &UnsignedLongObject{}
	case ast.Float32:
		return &FloatObject{}
	case ast.Float64:
		return &DoubleObject{}
	case ast.Char:
		return &CharacterObject{}
	case ast.String:
		return &StringObject{}
	}

	return nil
}
