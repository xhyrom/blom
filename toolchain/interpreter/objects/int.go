package objects

import (
	"blom/ast"
	"fmt"
)

type IntObject struct {
	value int32
}

func NewIntObject(value int32) *IntObject {
	return &IntObject{value: value}
}

func (i *IntObject) Type() ast.Type {
	return ast.Int32
}

func (i *IntObject) Inspect() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *IntObject) Value() interface{} {
	return i.value
}

func (i *IntObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = int32(value.(int8))
	case uint8:
		i.value = int32(value.(uint8))
	case int16:
		i.value = int32(value.(int16))
	case uint16:
		i.value = int32(value.(uint16))
	case int32:
		i.value = value.(int32)
	case uint32:
		i.value = int32(value.(uint32))
	case int64:
		i.value = int32(value.(int64))
	case uint64:
		i.value = int32(value.(uint64))
	case float32:
		i.value = int32(value.(float32))
	case float64:
		i.value = int32(value.(float64))
	}
}

func (i IntObject) Add(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value + o.value}
	}

	return nil
}

func (i *IntObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value - o.value}
	}

	return nil
}

func (i *IntObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value * o.value}
	}

	return nil
}

func (i *IntObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value / o.value}
	}

	return nil
}

func (i *IntObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value % o.value}
	}

	return nil
}

func (i *IntObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value & o.value}
	}

	return nil
}

func (i *IntObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value | o.value}
	}

	return nil
}

func (i *IntObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value ^ o.value}
	}

	return nil
}

func (i *IntObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value << uint(o.value)}
	}

	return nil
}

func (i *IntObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{value: i.value >> uint(o.value)}
	}

	return nil
}

func (i *IntObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *IntObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *IntObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *IntObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *IntObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *IntObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
