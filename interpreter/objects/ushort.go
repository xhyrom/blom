package objects

import (
	"blom/ast"
	"fmt"
)

type UnsignedShortObject struct {
	Value uint16
}

func (i *UnsignedShortObject) Type() ast.Type {
	return ast.UnsignedInt16
}

func (i *UnsignedShortObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i UnsignedShortObject) Add(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *UnsignedShortObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &UnsignedShortObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *UnsignedShortObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *UnsignedShortObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *UnsignedShortObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedShortObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
