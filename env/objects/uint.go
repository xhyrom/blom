package objects

import (
	"blom/compiler"
	"fmt"
)

type UnsignedIntObject struct {
	Value uint32
}

func (i *UnsignedIntObject) Type() compiler.Type {
	return compiler.UnsignedHalfword
}

func (i *UnsignedIntObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i UnsignedIntObject) Add(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *UnsignedIntObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &UnsignedIntObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *UnsignedIntObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *UnsignedIntObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *UnsignedIntObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *UnsignedIntObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
