package objects

import (
	"blom/ast"
	"fmt"
)

type UnsignedLongObject struct {
	value uint64
}

func (i *UnsignedLongObject) Type() ast.Type {
	return ast.UnsignedInt64
}

func (i *UnsignedLongObject) Inspect() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *UnsignedLongObject) Value() interface{} {
	return i.value
}

func (i *UnsignedLongObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = uint64(value.(int8))
	case uint8:
		i.value = uint64(value.(uint8))
	case int16:
		i.value = uint64(value.(int16))
	case uint16:
		i.value = uint64(value.(uint16))
	case int32:
		i.value = uint64(value.(int32))
	case uint32:
		i.value = uint64(value.(uint32))
	case int64:
		i.value = uint64(value.(int64))
	case uint64:
		i.value = value.(uint64)
	case float32:
		i.value = uint64(value.(float32))
	case float64:
		i.value = uint64(value.(float64))
	}
}

func (i UnsignedLongObject) Add(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value + o.value}
	}

	return nil
}

func (i *UnsignedLongObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value - o.value}
	}

	return nil
}

func (i *UnsignedLongObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value * o.value}
	}

	return nil
}

func (i *UnsignedLongObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value / o.value}
	}

	return nil
}

func (i *UnsignedLongObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value % o.value}
	}

	return nil
}

func (i *UnsignedLongObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value & o.value}
	}

	return nil
}

func (i *UnsignedLongObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value | o.value}
	}

	return nil
}

func (i *UnsignedLongObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value ^ o.value}
	}

	return nil
}

func (i *UnsignedLongObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value << uint(o.value)}
	}

	return nil
}

func (i *UnsignedLongObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{value: i.value >> uint(o.value)}
	}

	return nil
}

func (i *UnsignedLongObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *UnsignedLongObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *UnsignedLongObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *UnsignedLongObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *UnsignedLongObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
