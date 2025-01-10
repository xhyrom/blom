package objects

import (
	"blom/ast"
	"fmt"
)

type UnsignedShortObject struct {
	value uint16
}

func (i *UnsignedShortObject) Type() ast.Type {
	return ast.UnsignedInt16
}

func (i *UnsignedShortObject) Inspect() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *UnsignedShortObject) Value() interface{} {
	return i.value
}

func (i *UnsignedShortObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = uint16(value.(int8))
	case uint8:
		i.value = uint16(value.(uint8))
	case int16:
		i.value = uint16(value.(int16))
	case uint16:
		i.value = value.(uint16)
	case int32:
		i.value = uint16(value.(int32))
	case uint32:
		i.value = uint16(value.(uint32))
	case int64:
		i.value = uint16(value.(int64))
	case uint64:
		i.value = uint16(value.(uint64))
	case float32:
		i.value = uint16(value.(float32))
	case float64:
		i.value = uint16(value.(float64))
	}
}

func (i UnsignedShortObject) Add(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value + o.value}
	}

	return nil
}

func (i *UnsignedShortObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value - o.value}
	}

	return nil
}

func (i *UnsignedShortObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value * o.value}
	}

	return nil
}

func (i *UnsignedShortObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value / o.value}
	}

	return nil
}

func (i *UnsignedShortObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value % o.value}
	}

	return nil
}

func (i *UnsignedShortObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value & o.value}
	}

	return nil
}

func (i *UnsignedShortObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value | o.value}
	}

	return nil
}

func (i *UnsignedShortObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value ^ o.value}
	}

	return nil
}

func (i *UnsignedShortObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value << uint(o.value)}
	}

	return nil
}

func (i *UnsignedShortObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{value: i.value >> uint(o.value)}
	}

	return nil
}

func (i *UnsignedShortObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *UnsignedShortObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *UnsignedShortObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *UnsignedShortObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *UnsignedShortObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *UnsignedShortObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
