package objects

import (
	"blom/compiler"
	"fmt"
)

type UnsignedLongObject struct {
	Value uint64
}

func (i *UnsignedLongObject) Type() compiler.Type {
	return compiler.UnsignedLong
}

func (i *UnsignedLongObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i UnsignedLongObject) Add(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *UnsignedLongObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &UnsignedLongObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *UnsignedLongObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *UnsignedLongObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *UnsignedLongObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedLongObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
