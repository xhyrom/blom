package objects

import (
	"blom/ast"
	"fmt"
)

type LongObject struct {
	value int64
}

func NewLongObject(value int64) *LongObject {
	return &LongObject{value: value}
}

func (i *LongObject) Type() ast.Type {
	return ast.Int64
}

func (i *LongObject) Inspect() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *LongObject) Value() interface{} {
	return i.value
}

func (i *LongObject) SetValue(value interface{}) {
	switch value.(type) {
	case int8:
		i.value = int64(value.(int8))
	case uint8:
		i.value = int64(value.(uint8))
	case int16:
		i.value = int64(value.(int16))
	case uint16:
		i.value = int64(value.(uint16))
	case int32:
		i.value = int64(value.(int32))
	case uint32:
		i.value = int64(value.(uint32))
	case int64:
		i.value = value.(int64)
	case uint64:
		i.value = int64(value.(uint64))
	case float32:
		i.value = int64(value.(float32))
	case float64:
		i.value = int64(value.(float64))
	}
}

func (i LongObject) Add(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value + o.value}
	}

	return nil
}

func (i *LongObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value - o.value}
	}

	return nil
}

func (i *LongObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value * o.value}
	}

	return nil
}

func (i *LongObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value / o.value}
	}

	return nil
}

func (i *LongObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value % o.value}
	}

	return nil
}

func (i *LongObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value & o.value}
	}

	return nil
}

func (i *LongObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value | o.value}
	}

	return nil
}

func (i *LongObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value ^ o.value}
	}

	return nil
}

func (i *LongObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value << uint(o.value)}
	}

	return nil
}

func (i *LongObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{value: i.value >> uint(o.value)}
	}

	return nil
}

func (i *LongObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{value: i.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (i *LongObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{value: i.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (i *LongObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{value: i.value < o.value}
	}

	return nil
}

func (i *LongObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{value: i.value <= o.value}
	}

	return nil
}

func (i *LongObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{value: i.value > o.value}
	}

	return nil
}

func (i *LongObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{value: i.value >= o.value}
	}

	return nil
}
