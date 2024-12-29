package compiler

import (
	"errors"
	"fmt"
	"slices"
)

type Type int

const (
	UnsignedByte Type = iota
	UnsignedHalfword
	UnsignedWord
	UnsignedLong
	Byte
	Halfword
	Boolean
	Word
	Long
	Single
	Double
	Char
	String
	Void
	Null
)

var humanTypes = []string{
	UnsignedByte:     "u8",
	UnsignedHalfword: "u16",
	UnsignedWord:     "u32",
	UnsignedLong:     "u64",
	Byte:             "i8",
	Halfword:         "i16",
	Boolean:          "bool",
	Word:             "i32",
	Long:             "i64",
	Single:           "f32",
	Double:           "f64",
	Char:             "char",
	String:           "string",
	Void:             "void",
	Null:             "null",
}

var types = []string{
	UnsignedByte:     "ub",
	UnsignedHalfword: "uh",
	UnsignedWord:     "uw",
	UnsignedLong:     "ul",
	Byte:             "b",
	Halfword:         "h",
	Boolean:          "w",
	Word:             "w",
	Long:             "l",
	Single:           "s",
	Double:           "d",
	Char:             "c",
	String:           "l",
	Void:             "",
	Null:             "",
}

func ParseType(str string) (Type, error) {
	index := slices.Index(humanTypes, str)
	if index == -1 {
		return 0, errors.New(fmt.Sprintf("Unknown type '%s'", str))
	}

	return Type(index), nil
}

func (t Type) Inspect() string {
	return humanTypes[t]
}

func (t Type) String() string {
	return types[t]
}

func (t Type) IsNumeric() bool {
	return t == UnsignedByte || t == UnsignedHalfword || t == UnsignedWord || t == UnsignedLong || t == Byte || t == Halfword || t == Word || t == Long || t == Single || t == Double
}

func (t Type) IsInteger() bool {
	return t == UnsignedByte || t == UnsignedHalfword || t == UnsignedWord || t == UnsignedLong || t == Byte || t == Halfword || t == Word || t == Long
}

func (t Type) IsFloatingPoint() bool {
	return t == Single || t == Double
}
