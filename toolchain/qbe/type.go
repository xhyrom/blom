package qbe

import (
	"blom/ast"
	"fmt"
	"strings"
	"unsafe"
)

type Type interface {
	String() string
	IsNumeric() bool
	IsInteger() bool
	IsFloatingPoint() bool
	IsSigned() bool
	IsUnsigned() bool
	IsPointer() bool
	IsMapToInt() bool
	Weight() uint8
	Size() uint64
	IntoAbi() Type
}

type TypeId int

const (
	Byte TypeId = iota
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
	Pointer
	Char
	Boolean
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
	Pointer: "l",
	Boolean: "w",
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
		return PointerBox{Inner: Char}
	case ast.Void:
		return Void
	case ast.Null:
		return Null
	}

	if t.IsPointer() {
		return PointerBox{Inner: RemapAstType(t.Dereference())}
	}

	panic(fmt.Sprintf("Unknown type '%s'", t))
}

func (t TypeId) String() string {
	return types[t]
}

func (t TypeId) IsNumeric() bool {
	return t.IsInteger() || t.IsFloatingPoint()
}

func (t TypeId) IsInteger() bool {
	return t == UnsignedByte || t == UnsignedHalfword || t == UnsignedWord || t == UnsignedLong || t == Byte || t == Halfword || t == Word || t == Long
}

func (t TypeId) IsFloatingPoint() bool {
	return t == Single || t == Double
}

func (t TypeId) IsSigned() bool {
	return t == Byte || t == Halfword || t == Word || t == Long
}

func (t TypeId) IsUnsigned() bool {
	return t == UnsignedByte || t == UnsignedHalfword || t == UnsignedWord || t == UnsignedLong
}

func (t TypeId) IsPointer() bool {
	return t == Pointer
}

func (t TypeId) IsMapToInt() bool {
	switch t {
	case Byte, UnsignedByte, Halfword, UnsignedHalfword, UnsignedWord, Boolean, Char, Void:
		return true
	}

	return false
}

func (t TypeId) Weight() uint8 {
	switch t {
	case Double:
		return 4
	case Single:
		return 3
	case Long, UnsignedLong, Pointer:
		return 2
	case Word:
		return 1
	default:
		if t.IsMapToInt() {
			return 1
		}

		return 0
	}
}

// Return the size of a type in bytes
func (t TypeId) Size() uint64 {
	switch t {
	case UnsignedByte, Byte, Char:
		return 1
	case UnsignedHalfword, Halfword:
		return 2
	case Boolean, UnsignedWord, Word, Single, Void:
		return 4
	case Double:
		return 8
	// Returns 8 on 64-bit systems and 4 on 32-bit systems
	case UnsignedLong, Long, Pointer:
		return uint64(unsafe.Sizeof(uintptr(0)))
	default:
		panic(fmt.Sprintf("Unknown type '%s'", t))
	}
}

// Return a C ABI type
func (t TypeId) IntoAbi() Type {
	switch t {
	case Byte, Char, UnsignedByte, Halfword, UnsignedHalfword, UnsignedWord:
		return Word
	}

	return t
}

// QBE aggregate
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
