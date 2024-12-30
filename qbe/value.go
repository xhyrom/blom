package qbe

import "fmt"

type ValueType int

const (
	TemporaryValueType ValueType = iota
	GlobalValueType
	ConstantValueType
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
type ConstantValue struct {
	Prefix string
	Value  int64
}

func NewConstantValue(value int64) ConstantValue {
	return ConstantValue{Value: value}
}

func NewConstantValueWithPrefix(prefix string, value int64) ConstantValue {
	return ConstantValue{Prefix: prefix, Value: value}
}

func (v ConstantValue) Type() ValueType {
	return ConstantValueType
}

func (v ConstantValue) String() string {
	return fmt.Sprintf("%s%d", v.Prefix, v.Value)
}
