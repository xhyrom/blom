package ast

import (
	"errors"
	"fmt"
	"slices"
)

type Type int

const (
	Int8 Type = iota
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
}

func ParseType(str string) (Type, error) {
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

	return Type(index), nil
}

func NewPointerType(baseType Type) Type {
	return Type(int(Pointer) + int(baseType))
}

func (t Type) IsPointer() bool {
	return t >= Pointer
}

func (t Type) Dereference() Type {
	if !t.IsPointer() {
		panic(fmt.Sprintf("Type '%s' is not a pointer", t))
	}
	return Type(int(t) - int(Pointer))
}

func (t Type) String() string {
	if t.IsPointer() {
		return t.Dereference().String() + "*"
	}
	return types[t]
}

func (t Type) IsNumeric() bool {
	return t >= Int8 && t <= Float64
}

func (t Type) IsInteger() bool {
	return t >= Int8 && t <= UnsignedInt64
}

func (t Type) IsFloatingPoint() bool {
	return t == Float32 || t == Float64
}

func (t Type) IsMapToInt() bool {
	switch t {
	case Int8, UnsignedInt8, Int16, UnsignedInt16, UnsignedInt32, Boolean, Char, Void:
		return true
	}

	return false
}

func (t Type) Weight() uint8 {
	switch t {
	case Float64:
		return 4
	case Float32:
		return 3
	case Int64, UnsignedInt64, String:
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
