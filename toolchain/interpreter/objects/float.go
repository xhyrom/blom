package objects

import (
	"blom/ast"
	"fmt"
)

type FloatObject struct {
	value float32
}

func NewFloatObject(value float32) *FloatObject {
	return &FloatObject{value: value}
}

func (i *FloatObject) Type() ast.Type {
	return ast.Float32
}

func (i *FloatObject) Inspect() string {
	return fmt.Sprintf("%f", i.value)
}

func (i *FloatObject) Value() interface{} {
	return i.value
}

func (i *FloatObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = float32(value.(int8))
	case uint8:
		i.value = float32(value.(uint8))
	case int16:
		i.value = float32(value.(int16))
	case uint16:
		i.value = float32(value.(uint16))
	case int32:
		i.value = float32(value.(int32))
	case uint32:
		i.value = float32(value.(uint32))
	case int64:
		i.value = float32(value.(int64))
	case uint64:
		i.value = float32(value.(uint64))
	case float32:
		i.value = value.(float32)
	case float64:
		i.value = float32(value.(float64))
	}
}

func (i FloatObject) Add(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &FloatObject{value: i.value + o.value}
	}

	return nil
}

func (i *FloatObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &FloatObject{value: i.value - o.value}
	}

	return nil
}

func (i *FloatObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &FloatObject{value: i.value * o.value}
	}

	return nil
}

func (i *FloatObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &FloatObject{value: i.value / o.value}
	}

	return nil
}

func (i *FloatObject) Modulo(other Object) Object {
	return nil
}

func (i *FloatObject) BitwiseAnd(other Object) Object {
	return nil
}

func (i *FloatObject) BitwiseOr(other Object) Object {
	return nil
}

func (i *FloatObject) BitwiseXor(other Object) Object {
	return nil
}

func (i *FloatObject) LeftShift(other Object) Object {
	return nil
}

func (i *FloatObject) RightShift(other Object) Object {
	return nil
}

func (i *FloatObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *FloatObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *FloatObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return &BooleanObject{value: false}
}

func (i *FloatObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return &BooleanObject{value: false}
}

func (i *FloatObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return &BooleanObject{value: false}
}

func (i *FloatObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return &BooleanObject{value: false}
}
