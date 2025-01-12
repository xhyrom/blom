package objects

import (
	"blom/ast"
	"fmt"
)

type DoubleObject struct {
	value float64
}

func NewDoubleObject(value float64) *DoubleObject {
	return &DoubleObject{value: value}
}

func (i *DoubleObject) Type() ast.Type {
	return ast.Float64
}

func (i *DoubleObject) Inspect() string {
	return fmt.Sprintf("%f", i.value)
}

func (i *DoubleObject) Value() interface{} {
	return i.value
}

func (i *DoubleObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = float64(value.(int8))
	case uint8:
		i.value = float64(value.(uint8))
	case int16:
		i.value = float64(value.(int16))
	case uint16:
		i.value = float64(value.(uint16))
	case int32:
		i.value = float64(value.(int32))
	case uint32:
		i.value = float64(value.(uint32))
	case int64:
		i.value = float64(value.(int64))
	case uint64:
		i.value = float64(value.(uint64))
	case float32:
		i.value = float64(value.(float32))
	case float64:
		i.value = value.(float64)
	}
}

func (i DoubleObject) Add(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &DoubleObject{value: i.value + o.value}
	}

	return nil
}

func (i *DoubleObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &DoubleObject{value: i.value - o.value}
	}

	return nil
}

func (i *DoubleObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &DoubleObject{value: i.value * o.value}
	}

	return nil
}

func (i *DoubleObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &DoubleObject{value: i.value / o.value}
	}

	return nil
}

func (i *DoubleObject) Modulo(other Object) Object {
	return nil
}

func (i *DoubleObject) BitwiseAnd(other Object) Object {
	return nil
}

func (i *DoubleObject) BitwiseOr(other Object) Object {
	return nil
}

func (i *DoubleObject) BitwiseXor(other Object) Object {
	return nil
}

func (i *DoubleObject) LeftShift(other Object) Object {
	return nil
}

func (i *DoubleObject) RightShift(other Object) Object {
	return nil
}

func (i *DoubleObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *DoubleObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *DoubleObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *DoubleObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *DoubleObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *DoubleObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
