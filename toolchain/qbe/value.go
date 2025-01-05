package qbe

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type ValueType int

const (
	TemporaryValueType ValueType = iota
	GlobalValueType
	ConstantValueType
	LiteralValueType
)

type Value interface {
	Type() ValueType
	String() string
}

// TemporaryValue represents a temporary value. (%name)
type TemporaryValue struct {
	Name string
}

func NewTemporaryValue(name string) TemporaryValue {
	return TemporaryValue{Name: name}
}

func (v TemporaryValue) Type() ValueType {
	return TemporaryValueType
}

func (v TemporaryValue) String() string {
	return fmt.Sprintf("%%%s", v.Name)
}

// GlobalValue represents a global value. ($name)
type GlobalValue struct {
	Name string
}

func NewGlobalValue(name string) GlobalValue {
	return GlobalValue{Name: name}
}

func (v GlobalValue) Type() ValueType {
	return GlobalValueType
}

func (v GlobalValue) String() string {
	return fmt.Sprintf("$%s", v.Name)
}

type TypedValue struct {
	Value Value
	Type  Type
}

func NewTypedValue(t Type, value Value) TypedValue {
	return TypedValue{Value: value, Type: t}
}

func (v TypedValue) String() string {
	return fmt.Sprintf("%s %s", v.Type, v.Value)
}

func (v TypedValue) AbiString() string {
	return fmt.Sprintf("%s %s", v.Type.IntoAbi(), v.Value)
}

// ConstantValue represents a constant value. (prefix value)
// For example, an integer value would be represented as "42".
// A 64-bit floating point value would be represented as "d_42".
// A 32-bit floating point value would be represented as "s_42".
type ConstantValue[T constraints.Integer | constraints.Float] struct {
	Prefix string
	Value  T
}

func NewConstantValue[T constraints.Integer | constraints.Float](value T) ConstantValue[T] {
	return ConstantValue[T]{Value: value}
}

func NewConstantValueWithPrefix[T constraints.Integer | constraints.Float](prefix string, value T) ConstantValue[T] {
	return ConstantValue[T]{Prefix: prefix, Value: value}
}

func (v ConstantValue[T]) Type() ValueType {
	return ConstantValueType
}

func (v ConstantValue[T]) String() string {
	return fmt.Sprintf("%s%v", v.Prefix, v.Value)
}

// LiteralValue represents a literal value. (prefix value)
type LiteralValue struct {
	Prefix string
	Value  string
}

func NewLiteralValue(value string) LiteralValue {
	return LiteralValue{Value: value}
}

func NewLiteralValueWithPrefix(prefix string, value string) LiteralValue {
	return LiteralValue{Prefix: prefix, Value: value}
}

func (v LiteralValue) Type() ValueType {
	return LiteralValueType
}

func (v LiteralValue) String() string {
	return fmt.Sprintf("%s%s", v.Prefix, v.Value)
}
