package objects

import (
	"blom/ast"
	"fmt"
)

type UnsignedIntObject struct {
	value uint32
}

func NewUnsignedIntObject(value uint32) *UnsignedIntObject {
	return &UnsignedIntObject{value: value}
}

func (i *UnsignedIntObject) Type() ast.Type {
	return ast.UnsignedInt32
}

func (i *UnsignedIntObject) Inspect() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *UnsignedIntObject) Value() interface{} {
	return i.value
}

func (i *UnsignedIntObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = uint32(value.(int8))
	case uint8:
		i.value = uint32(value.(uint8))
	case int16:
		i.value = uint32(value.(int16))
	case uint16:
		i.value = uint32(value.(uint16))
	case int32:
		i.value = uint32(value.(int32))
	case uint32:
		i.value = value.(uint32)
	case int64:
		i.value = uint32(value.(int64))
	case uint64:
		i.value = uint32(value.(uint64))
	case float32:
		i.value = uint32(value.(float32))
	case float64:
		i.value = uint32(value.(float64))
	}
}

func (i UnsignedIntObject) Add(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value + o.value}
	}

	return nil
}

func (i *UnsignedIntObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value - o.value}
	}

	return nil
}

func (i *UnsignedIntObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value * o.value}
	}

	return nil
}

func (i *UnsignedIntObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value / o.value}
	}

	return nil
}

func (i *UnsignedIntObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value % o.value}
	}

	return nil
}

func (i *UnsignedIntObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value & o.value}
	}

	return nil
}

func (i *UnsignedIntObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value | o.value}
	}

	return nil
}

func (i *UnsignedIntObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value ^ o.value}
	}

	return nil
}

func (i *UnsignedIntObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value << uint(o.value)}
	}

	return nil
}

func (i *UnsignedIntObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{value: i.value >> uint(o.value)}
	}

	return nil
}

func (i *UnsignedIntObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *UnsignedIntObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *UnsignedIntObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *UnsignedIntObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *UnsignedIntObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *UnsignedIntObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
