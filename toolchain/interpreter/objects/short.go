package objects

import (
	"blom/ast"
	"fmt"
)

type ShortObject struct {
	value int16
}

func NewShortObject(value int16) *ShortObject {
	return &ShortObject{value: value}
}

func (i *ShortObject) Type() ast.Type {
	return ast.Int16
}

func (i *ShortObject) Inspect() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *ShortObject) Value() interface{} {
	return i.value
}

func (i *ShortObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = int16(value.(int8))
	case uint8:
		i.value = int16(value.(uint8))
	case int16:
		i.value = value.(int16)
	case uint16:
		i.value = int16(value.(uint16))
	case int32:
		i.value = int16(value.(int32))
	case uint32:
		i.value = int16(value.(uint32))
	case int64:
		i.value = int16(value.(int64))
	case uint64:
		i.value = int16(value.(uint64))
	case float32:
		i.value = int16(value.(float32))
	case float64:
		i.value = int16(value.(float64))
	}
}

func (i ShortObject) Add(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value + o.value}
	}

	return nil
}

func (i *ShortObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value - o.value}
	}

	return nil
}

func (i *ShortObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value * o.value}
	}

	return nil
}

func (i *ShortObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value / o.value}
	}

	return nil
}

func (i *ShortObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value % o.value}
	}

	return nil
}

func (i *ShortObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value & o.value}
	}

	return nil
}

func (i *ShortObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value | o.value}
	}

	return nil
}

func (i *ShortObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value ^ o.value}
	}

	return nil
}

func (i *ShortObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value << uint(o.value)}
	}

	return nil
}

func (i *ShortObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{value: i.value >> uint(o.value)}
	}

	return nil
}

func (i *ShortObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *ShortObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *ShortObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *ShortObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *ShortObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *ShortObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
