package qbe

import (
	"unsafe"
)

// PointerBox is a wrapper around a Type that represents a pointer.
// It holds a reference to the inner Type.
type PointerBox struct {
	Inner Type
}

func NewPointer(inner Type) PointerBox {
	return PointerBox{Inner: inner}
}

func (p PointerBox) String() string {
	return Pointer.String()
}

func (p PointerBox) IsNumeric() bool {
	return p.Inner.IsNumeric()
}

func (p PointerBox) IsInteger() bool {
	return p.Inner.IsInteger()
}

func (p PointerBox) IsFloatingPoint() bool {
	return p.Inner.IsFloatingPoint()
}

func (p PointerBox) IsSigned() bool {
	return p.Inner.IsSigned()
}

func (p PointerBox) IsUnsigned() bool {
	return p.Inner.IsUnsigned()
}

func (p PointerBox) IsPointer() bool {
	return true
}

func (p PointerBox) IsMapToInt() bool {
	return Pointer.IsMapToInt()
}

func (p PointerBox) Weight() uint8 {
	return Pointer.Weight()
}

func (p PointerBox) Size() uint64 {
	return uint64(unsafe.Sizeof(uintptr(0)))
}

func (p PointerBox) IntoAbi() Type {
	return Pointer
}
