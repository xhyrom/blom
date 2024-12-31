package objects

import (
	"blom/ast"
	"fmt"
)

type DoubleObject struct {
	Value float64
}

func (i *DoubleObject) Type() ast.Type {
	return ast.Float64
}

func (i *DoubleObject) Inspect() string {
	return fmt.Sprintf("%f", i.Value)
}

func (i DoubleObject) Add(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &DoubleObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *DoubleObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &DoubleObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *DoubleObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &DoubleObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *DoubleObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &DoubleObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *DoubleObject) Modulo(other Object) Object {
	return nil
}

func (i *DoubleObject) BitwiseAnd(other Object) Object {
	return nil
}

func (i *DoubleObject) BitwiseOr(other Object) Object {
	return nil
}

func (i *DoubleObject) BitwiseXor(other Object) Object {
	return nil
}

func (i *DoubleObject) LeftShift(other Object) Object {
	return nil
}

func (i *DoubleObject) RightShift(other Object) Object {
	return nil
}

func (i *DoubleObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *DoubleObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return nil
}

func (i *DoubleObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return nil
}

func (i *DoubleObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return nil
}

func (i *DoubleObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *DoubleObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return nil
}
