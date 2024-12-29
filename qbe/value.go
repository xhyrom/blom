package qbe

import "fmt"

type ValueType int

const (
	TemporaryValueType ValueType = iota
	GlobalValueType
)

type Value interface {
	Type() ValueType
	String() string
}

// TemporaryValue represents a temporary value. (%name)
type TemporaryValue struct {
	Name  string
	Value string
}

func (v TemporaryValue) Type() ValueType {
	return TemporaryValueType
}

func (v TemporaryValue) String() string {
	return fmt.Sprintf("%%%s", v.Name)
}

// GlobalValue represents a global value. ($name)
type GlobalValue struct {
	Name  string
	Value string
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

func (v TypedValue) String() string {
	return fmt.Sprintf("%s %s", v.Type, v.Value)
}

func (v TypedValue) AbiString() string {
	return fmt.Sprintf("%s %s", v.Type.IntoAbi(), v.Value)
}
