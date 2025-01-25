package ast

import (
	"fmt"
	"slices"
)

type Type interface {
	Equal(other Type) bool
	String() string
	IsPointer() bool
	IsFunction() bool
	IsNumeric() bool
	IsInteger() bool
	IsFloatingPoint() bool
	IsMapToInt() bool
	Weight() uint8
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

func ParseType(str string, additionalTypes map[string]Type) (Type, error) {
	if len(str) > 1 && str[len(str)-1] == '*' {
		baseType, err := ParseType(str[:len(str)-1], additionalTypes)
		if err != nil {
			return nil, err
		}
		return NewPointer(baseType), nil
	}

	index := slices.Index(types, str)
	if index != -1 {
		return TypeId(index), nil
	}

	if t, exists := additionalTypes[str]; exists {
		return t, nil
	}

	return nil, fmt.Errorf("unknown type '%s'", str)
}

func (t TypeId) Equal(other Type) bool {
	return t == other
}

func (t TypeId) String() string {
	return types[t]
}

func (t TypeId) IsPointer() bool {
	return t == Pointer
}

func (t TypeId) IsFunction() bool {
	return t == Function
}

func (t TypeId) IsNumeric() bool {
	return t.IsInteger() || t.IsFloatingPoint()
}

func (t TypeId) IsInteger() bool {
	return t == Int8 || t == UnsignedInt8 || t == Int16 || t == UnsignedInt16 || t == Int32 || t == UnsignedInt32 || t == Int64 || t == UnsignedInt64
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

// PointerType is a wrapper around a Type that represents a pointer.
// It holds a reference to the inner Type.
type PointerType struct {
	Inner Type
}

func NewPointer(inner Type) PointerType {
	return PointerType{Inner: inner}
}

func (p PointerType) Equal(other Type) bool {
	if otherPointer, ok := other.(PointerType); ok {
		return p.Inner == otherPointer.Inner
	}

	return false
}

func (p PointerType) String() string {
	return fmt.Sprintf("*%s", p.Inner.String())
}

func (p PointerType) IsNumeric() bool {
	return p.Inner.IsNumeric()
}

func (p PointerType) IsInteger() bool {
	return p.Inner.IsInteger()
}

func (p PointerType) IsFloatingPoint() bool {
	return p.Inner.IsFloatingPoint()
}

func (p PointerType) IsPointer() bool {
	return true
}

func (p PointerType) IsVoidPointer() bool {
	return p.Inner == Void
}

func (p PointerType) IsFunction() bool {
	return p.Inner.IsFunction()
}

func (p PointerType) IsMapToInt() bool {
	return Pointer.IsMapToInt()
}

func (p PointerType) Weight() uint8 {
	return Pointer.Weight()
}

func (p PointerType) Dereference() Type {
	return p.Inner
}

// FunctionType is a wrapper around a Type that represents a function.
// It holds a reference to the inner Type.
type FunctionType struct {
	Arguments  []Type
	ReturnType Type
}

func NewFunctionType(args []Type, returnType Type) FunctionType {
	return FunctionType{Arguments: args, ReturnType: returnType}
}

func (f FunctionType) Equal(other Type) bool {
	if otherFunction, ok := other.(FunctionType); ok {
		if len(f.Arguments) != len(otherFunction.Arguments) {
			return false
		}

		for i, arg := range f.Arguments {
			if !arg.Equal(otherFunction.Arguments[i]) {
				return false
			}
		}

		return f.ReturnType.Equal(otherFunction.ReturnType)
	}

	return false
}

func (f FunctionType) String() string {
	return Function.String()
}

func (f FunctionType) IsPointer() bool {
	return Function.IsPointer()
}

func (f FunctionType) IsFunction() bool {
	return true
}

func (f FunctionType) IsNumeric() bool {
	return Function.IsNumeric()
}

func (f FunctionType) IsInteger() bool {
	return Function.IsInteger()
}

func (f FunctionType) IsFloatingPoint() bool {
	return Function.IsFloatingPoint()
}

func (f FunctionType) IsMapToInt() bool {
	return Function.IsMapToInt()
}

func (f FunctionType) Weight() uint8 {
	return Function.Weight()
}

func (f FunctionType) AsId() TypeId {
	return Function
}
