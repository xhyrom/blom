package qbe

import (
	"blom/ast"
	"fmt"
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

func RemapAstType(t ast.Type) Type {
	switch t {
	case ast.Int8:
		return Byte
	case ast.UnsignedInt8:
		return UnsignedByte
	case ast.Int16:
		return Halfword
	case ast.UnsignedInt16:
		return UnsignedHalfword
	case ast.Int32:
		return Word
	case ast.UnsignedInt32:
		return UnsignedWord
	case ast.Int64:
		return Long
	case ast.UnsignedInt64:
		return UnsignedLong
	case ast.Float32:
		return Single
	case ast.Float64:
		return Double
	case ast.Boolean:
		return Boolean
	case ast.Char:
		return Char
	case ast.String:
		return String
	case ast.Void:
		return Void
	case ast.Null:
		return Null
	}

	panic(fmt.Sprintf("Unknown type '%s'", t))
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
