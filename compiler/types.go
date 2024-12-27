package compiler

import (
	"errors"
	"fmt"
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
	Boolean:          "bool",
	Word:             "w",
	Long:             "l",
	Single:           "f",
	Double:           "d",
	Char:             "c",
	Void:             "",
	Null:             "",
}

var mapping = map[string]Type{
	"i8":   Byte,
	"i16":  Halfword,
	"i32":  Word,
	"i64":  Long,
	"u8":   UnsignedByte,
	"u16":  UnsignedHalfword,
	"u32":  UnsignedWord,
	"u64":  UnsignedLong,
	"f32":  Single,
	"f64":  Double,
	"char": Char,
	"void": Void,
}

func ParseType(str string) (Type, error) {
	if val, ok := mapping[str]; ok {
		return val, nil
	}

	return 0, errors.New(fmt.Sprintf("Unknown type \"%s\"", str))
}

func (t Type) Inspect() string {
	return humanTypes[t]
}

func (t Type) String() string {
	return types[t]
}
