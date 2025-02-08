package qbe

import (
	"fmt"
)

type StructBox struct {
	Name string
}

func NewStruct(name string) StructBox {
	return StructBox{Name: name}
}

func (s StructBox) String() string {
	return fmt.Sprintf(":%s", s.Name)
}

func (s StructBox) IsNumeric() bool {
	return Struct.IsNumeric()
}

func (s StructBox) IsInteger() bool {
	return Struct.IsInteger()
}

func (s StructBox) IsFloatingPoint() bool {
	return Struct.IsFloatingPoint()
}

func (s StructBox) IsSigned() bool {
	return Struct.IsSigned()
}

func (s StructBox) IsUnsigned() bool {
	return Struct.IsUnsigned()
}

func (s StructBox) IsPointer() bool {
	return Struct.IsPointer()
}

func (s StructBox) IsFunction() bool {
	return Struct.IsFunction()
}

func (s StructBox) IsStruct() bool {
	return Struct.IsStruct()
}

func (s StructBox) IsMapToInt() bool {
	return Struct.IsMapToInt()
}

func (s StructBox) Weight() uint8 {
	return Struct.Weight()
}

func (s StructBox) Size(module Module) uint64 {
	return module.GetTypeByName(s.Name).Size(module)
}

func (s StructBox) IntoAbi() Type {
	return s
}

func (s StructBox) IntoBase() Type {
	return Struct.IntoBase()
}
