package objects

import (
	"blom/ast"
	"fmt"
)

type UnsignedByteObject struct {
	value uint8
}

func (i *UnsignedByteObject) Type() ast.Type {
	return ast.UnsignedInt8
}

func (i *UnsignedByteObject) Inspect() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *UnsignedByteObject) Value() interface{} {
	return i.value
}

func (i *UnsignedByteObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = uint8(value.(int8))
	case uint8:
		i.value = value.(uint8)
	case int16:
		i.value = uint8(value.(int16))
	case uint16:
		i.value = uint8(value.(uint16))
	case int32:
		i.value = uint8(value.(int32))
	case uint32:
		i.value = uint8(value.(uint32))
	case int64:
		i.value = uint8(value.(int64))
	case uint64:
		i.value = uint8(value.(uint64))
	case float32:
		i.value = uint8(value.(float32))
	case float64:
		i.value = uint8(value.(float64))
	}
}

func (i UnsignedByteObject) Add(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value + o.value}
	}

	return nil
}

func (i *UnsignedByteObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value - o.value}
	}

	return nil
}

func (i *UnsignedByteObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value * o.value}
	}

	return nil
}

func (i *UnsignedByteObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value / o.value}
	}

	return nil
}

func (i *UnsignedByteObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value % o.value}
	}

	return nil
}

func (i *UnsignedByteObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value & o.value}
	}

	return nil
}

func (i *UnsignedByteObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value | o.value}
	}

	return nil
}

func (i *UnsignedByteObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value ^ o.value}
	}

	return nil
}

func (i *UnsignedByteObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value << uint(o.value)}
	}

	return nil
}

func (i *UnsignedByteObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{value: i.value >> uint(o.value)}
	}

	return nil
}

func (i *UnsignedByteObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *UnsignedByteObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *UnsignedByteObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *UnsignedByteObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *UnsignedByteObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *UnsignedByteObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
