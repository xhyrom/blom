package ast

import (
	"errors"
	"fmt"
	"slices"
)

type Type interface {
	IsPointer() bool
	IsFunction() bool
	Dereference() Type
	String() string
	IsNumeric() bool
	IsInteger() bool
	IsFloatingPoint() bool
	IsMapToInt() bool
	Weight() uint8
	AsId() TypeId
}

type TypeId int

const (
	Int8 TypeId = iota
	UnsignedInt8
	Int16
	UnsignedInt16
	Int32
	UnsignedInt32
	Int64
	UnsignedInt64
	Float32
	Float64
	Boolean
	Char
	String
	Void
	Null
	Function
	Pointer
)

var types = []string{
	Int8:          "i8",
	UnsignedInt8:  "u8",
	Int16:         "i16",
	UnsignedInt16: "u16",
	Int32:         "i32",
	UnsignedInt32: "u32",
	Int64:         "i64",
	UnsignedInt64: "u64",
	Float32:       "f32",
	Float64:       "f64",
	Boolean:       "bool",
	Char:          "char",
	String:        "string",
	Void:          "void",
	Null:          "null",
	Pointer:       "ptr",
	Function:      "fun",
}

func ParseType(str string) (TypeId, error) {
	if len(str) > 1 && str[len(str)-1] == '*' {
		baseType, err := ParseType(str[:len(str)-1])
		if err != nil {
			return -1, err
		}
		return NewPointerType(baseType), nil
	}

	index := slices.Index(types, str)
	if index == -1 {
		return -1, errors.New(fmt.Sprintf("Unknown type '%s'", str))
	}

	return TypeId(index), nil
}

func NewPointerType(baseType TypeId) TypeId {
	return TypeId(int(Pointer) + int(baseType))
}

func (t TypeId) IsPointer() bool {
	return t >= Pointer
}

func (t TypeId) IsFunction() bool {
	return t == Function
}

func (t TypeId) IsLambda() bool {
	return t == Function
}

func (t TypeId) Dereference() Type {
	if !t.IsPointer() {
		panic(fmt.Sprintf("Type '%s' is not a pointer", t))
	}
	return TypeId(int(t) - int(Pointer))
}

func (t TypeId) String() string {
	if t.IsPointer() {
		return t.Dereference().String() + "*"
	}
	return types[t]
}

func (t TypeId) IsNumeric() bool {
	return t >= Int8 && t <= Float64
}

func (t TypeId) IsInteger() bool {
	return t >= Int8 && t <= UnsignedInt64
}

func (t TypeId) IsFloatingPoint() bool {
	return t == Float32 || t == Float64
}

func (t TypeId) IsMapToInt() bool {
	switch t {
	case Int8, UnsignedInt8, Int16, UnsignedInt16, UnsignedInt32, Boolean, Char, Void:
		return true
	}

	return false
}

func (t TypeId) Weight() uint8 {
	switch t {
	case Float64:
		return 4
	case Float32:
		return 3
	case Int64, UnsignedInt64, String, Function:
		return 2
	case Int32:
		return 1
	default:
		if t.IsMapToInt() {
			return 1
		}

		return 0
	}
}

func (t TypeId) AsId() TypeId {
	return t
}

type FunctionBox struct {
	Inner LambdaDeclaration
}

func NewFunctionBox(inner LambdaDeclaration) FunctionBox {
	return FunctionBox{Inner: inner}
}

func (f FunctionBox) IsPointer() bool {
	return false
}

func (f FunctionBox) IsFunction() bool {
	return true
}

func (f FunctionBox) Dereference() Type {
	panic("FunctionBox is not a pointer")
}

func (f FunctionBox) String() string {
	return Function.String()
}

func (f FunctionBox) IsNumeric() bool {
	return Function.IsNumeric()
}

func (f FunctionBox) IsInteger() bool {
	return Function.IsInteger()
}

func (f FunctionBox) IsFloatingPoint() bool {
	return Function.IsFloatingPoint()
}

func (f FunctionBox) IsMapToInt() bool {
	return Function.IsMapToInt()
}

func (f FunctionBox) Weight() uint8 {
	return Function.Weight()
}

func (f FunctionBox) AsId() TypeId {
	return Function
}
