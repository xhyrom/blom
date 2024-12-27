package objects

import (
	"blom/compiler"
	"fmt"
)

type ByteObject struct {
	Value int8
}

func (i *ByteObject) Type() compiler.Type {
	return compiler.Byte
}

func (i *ByteObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i ByteObject) Add(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *ByteObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *ByteObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *ByteObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *ByteObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *ByteObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *ByteObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *ByteObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *ByteObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *ByteObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &ByteObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *ByteObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *ByteObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *ByteObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *ByteObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *ByteObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *ByteObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
