package objects

import (
	"blom/compiler"
	"fmt"
)

type IntObject struct {
	Value int32
}

func (i *IntObject) Type() compiler.Type {
	return compiler.Word
}

func (i *IntObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i IntObject) Add(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *IntObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *IntObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *IntObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *IntObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *IntObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *IntObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *IntObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *IntObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *IntObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &IntObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *IntObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *IntObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *IntObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *IntObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *IntObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *IntObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
