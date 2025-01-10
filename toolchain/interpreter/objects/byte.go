package objects

import (
	"blom/ast"
	"fmt"
)

type ByteObject struct {
	value int8
}

func (i *ByteObject) Type() ast.Type {
	return ast.Int8
}

func (i *ByteObject) Inspect() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *ByteObject) Value() interface{} {
	return i.value
}

func (i *ByteObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = value.(int8)
	case uint8:
		i.value = int8(value.(uint8))
	case int16:
		i.value = int8(value.(int16))
	case uint16:
		i.value = int8(value.(uint16))
	case int32:
		i.value = int8(value.(int32))
	case uint32:
		i.value = int8(value.(uint32))
	case int64:
		i.value = int8(value.(int64))
	case uint64:
		i.value = int8(value.(uint64))
	case float32:
		i.value = int8(value.(float32))
	case float64:
		i.value = int8(value.(float64))
	}
}

func (i ByteObject) Add(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value + o.value}
	}

	return nil
}

func (i *ByteObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value - o.value}
	}

	return nil
}

func (i *ByteObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value * o.value}
	}

	return nil
}

func (i *ByteObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value / o.value}
	}

	return nil
}

func (i *ByteObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value % o.value}
	}

	return nil
}

func (i *ByteObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value & o.value}
	}

	return nil
}

func (i *ByteObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value | o.value}
	}

	return nil
}

func (i *ByteObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value ^ o.value}
	}

	return nil
}

func (i *ByteObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value << uint(o.value)}
	}

	return nil
}

func (i *ByteObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{value: i.value >> uint(o.value)}
	}

	return nil
}

func (i *ByteObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *ByteObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *ByteObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *ByteObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *ByteObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *ByteObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
