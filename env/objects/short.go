package objects

import (
	"blom/compiler"
	"fmt"
)

type ShortObject struct {
	Value int16
}

func (i *ShortObject) Type() compiler.Type {
	return compiler.Halfword
}

func (i *ShortObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i ShortObject) Add(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *ShortObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *ShortObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *ShortObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *ShortObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *ShortObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *ShortObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *ShortObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *ShortObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *ShortObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &ShortObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *ShortObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *ShortObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *ShortObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *ShortObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *ShortObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *ShortObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
