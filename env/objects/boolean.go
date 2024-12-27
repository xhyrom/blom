package objects

import (
	"blom/compiler"
	"fmt"
)

type BooleanObject struct {
	Value bool
}

func (b *BooleanObject) Type() compiler.Type {
	return compiler.Boolean
}

func (b *BooleanObject) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
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
		return &BooleanObject{Value: b.Value == o.Value}
	}

	return &BooleanObject{Value: false}
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
