package qbe

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type Type int

const (
	Byte Type = iota
	UnsignedByte
	Halfword
	UnsignedHalfword
	Word
	UnsignedWord
	Long
	UnsignedLong
	Single
	Double
	// Custom
	Char
	Boolean
	String
	Void
	Null
)

var humanTypes = []string{
	Byte:             "i8",
	UnsignedByte:     "u8",
	Halfword:         "i16",
	UnsignedHalfword: "u16",
	Word:             "i32",
	UnsignedWord:     "u32",
	Long:             "i64",
	UnsignedLong:     "u64",
	Single:           "f32",
	Double:           "f64",
	// Custom
	Boolean: "bool",
	Char:    "char",
	String:  "string",
	Void:    "void",
	Null:    "null",
}

var types = []string{
	Byte:             "b",
	UnsignedByte:     "ub",
	Char:             "b",
	Halfword:         "h",
	UnsignedHalfword: "uh",
	Word:             "w",
	UnsignedWord:     "uw",
	Long:             "l",
	UnsignedLong:     "ul",
	Single:           "s",
	Double:           "d",
	// Custom
	Boolean: "w",
	String:  "l",
	Void:    "w",
	Null:    "",
}

func ParseHumanType(str string) (Type, error) {
	index := slices.Index(humanTypes, str)
	if index == -1 {
		return 0, errors.New(fmt.Sprintf("Unknown type '%s'", str))
	}

	return Type(index), nil
}

func (t Type) IntoHumanReadable() string {
	return humanTypes[t]
}

func (t Type) String() string {
	return types[t]
}

func (t Type) IsNumeric() bool {
	return t.IsInteger() || t.IsFloatingPoint()
}

func (t Type) IsInteger() bool {
	return t == UnsignedByte || t == UnsignedHalfword || t == UnsignedWord || t == UnsignedLong || t == Byte || t == Halfword || t == Word || t == Long
}

func (t Type) IsFloatingPoint() bool {
	return t == Single || t == Double
}

func (t Type) IsSigned() bool {
	return t == Byte || t == Halfword || t == Word || t == Long
}

func (t Type) IsUnsigned() bool {
	return t == UnsignedByte || t == UnsignedHalfword || t == UnsignedWord || t == UnsignedLong
}

// Return a C ABI type
func (t Type) IntoAbi() Type {
	switch t {
	case Byte, Char, UnsignedByte, Halfword, UnsignedHalfword, UnsignedWord:
		return Word
	}

	return t
}

type TypedTypeDefinitionItem struct {
	Count uint
	Type  Type
}

// QBE aggregate type definition
type TypeDefinition struct {
	Name  string
	Align *uint64
	Items []TypedTypeDefinitionItem
}

func (t TypeDefinition) String() string {
	result := fmt.Sprintf("type :%s = ", t.Name)

	if t.Align != nil {
		result += fmt.Sprintf("align %d ", *t.Align)
	}

	var parts []string
	for _, item := range t.Items {
		if item.Count > 1 {
			parts = append(parts, fmt.Sprintf("%s %d", item.Type, item.Count))
		} else {
			parts = append(parts, fmt.Sprintf("%s", item.Type))
		}
	}

	result += fmt.Sprintf("{ %s }", strings.Join(parts, ", "))

	return result
}
