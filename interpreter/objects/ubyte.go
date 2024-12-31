package objects

import (
	"blom/ast"
	"fmt"
)

type UnsignedByteObject struct {
	Value uint8
}

func (i *UnsignedByteObject) Type() ast.Type {
	return ast.UnsignedInt8
}

func (i *UnsignedByteObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i UnsignedByteObject) Add(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *UnsignedByteObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &UnsignedByteObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *UnsignedByteObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *UnsignedByteObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *UnsignedByteObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedByteObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
