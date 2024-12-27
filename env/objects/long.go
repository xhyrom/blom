package objects

import (
	"blom/compiler"
	"fmt"
)

type LongObject struct {
	Value int64
}

func (i *LongObject) Type() compiler.Type {
	return compiler.Long
}

func (i *LongObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i LongObject) Add(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *LongObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *LongObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *LongObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *LongObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *LongObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *LongObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *LongObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *LongObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *LongObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &LongObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *LongObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *LongObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *LongObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *LongObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *LongObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *LongObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
