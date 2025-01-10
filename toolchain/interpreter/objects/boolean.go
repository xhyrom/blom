package objects

import (
	"blom/ast"
	"fmt"
)

type BooleanObject struct {
	value bool
}

func NewBooleanObject(value bool) *BooleanObject {
	return &BooleanObject{value: value}
}

func (b *BooleanObject) Type() ast.Type {
	return ast.Boolean
}

func (b *BooleanObject) Inspect() string {
	return fmt.Sprintf("%t", b.value)
}

func (b *BooleanObject) Value() interface{} {
	return b.value
}

func (b *BooleanObject) SetValue(value interface{}) {
	v := value.(bool)
	b.value = v
}

func (b *BooleanObject) Add(other Object) Object {
	return nil
}

func (b *BooleanObject) Subtract(other Object) Object {
	return nil
}

func (b *BooleanObject) Multiply(other Object) Object {
	return nil
}

func (b *BooleanObject) Divide(other Object) Object {
	return nil
}

func (b *BooleanObject) Modulo(other Object) Object {
	return nil
}

func (b *BooleanObject) BitwiseAnd(other Object) Object {
	return nil
}

func (b *BooleanObject) BitwiseOr(other Object) Object {
	return nil
}

func (b *BooleanObject) BitwiseXor(other Object) Object {
	return nil
}

func (b *BooleanObject) LeftShift(other Object) Object {
	return nil
}

func (b *BooleanObject) RightShift(other Object) Object {
	return nil
}

func (b *BooleanObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *BooleanObject:
		return &BooleanObject{value: b.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (b *BooleanObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *BooleanObject:
		return &BooleanObject{value: b.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (b *BooleanObject) LessThan(other Object) Object {
	return nil
}

func (b *BooleanObject) LessThanOrEqual(other Object) Object {
	return nil
}

func (b *BooleanObject) GreaterThan(other Object) Object {
	return nil
}

func (b *BooleanObject) GreaterThanOrEqual(other Object) Object {
	return nil
}
