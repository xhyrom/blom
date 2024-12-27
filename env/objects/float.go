package objects

import (
	"blom/compiler"
	"fmt"
)

type FloatObject struct {
	Value float32
}

func (i *FloatObject) Type() compiler.Type {
	return compiler.Single
}

func (i *FloatObject) Inspect() string {
	return fmt.Sprintf("%f", i.Value)
}

func (i FloatObject) Add(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &FloatObject{Value: i.Value + o.Value}
	}

	return nil
}

func (i *FloatObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &FloatObject{Value: i.Value - o.Value}
	}

	return nil
}

func (i *FloatObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &FloatObject{Value: i.Value * o.Value}
	}

	return nil
}

func (i *FloatObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &FloatObject{Value: i.Value / o.Value}
	}

	return nil
}

func (i *FloatObject) Modulo(other Object) Object {
	return nil
}

func (i *FloatObject) BitwiseAnd(other Object) Object {
	return nil
}

func (i *FloatObject) BitwiseOr(other Object) Object {
	return nil
}

func (i *FloatObject) BitwiseXor(other Object) Object {
	return nil
}

func (i *FloatObject) LeftShift(other Object) Object {
	return nil
}

func (i *FloatObject) RightShift(other Object) Object {
	return nil
}

func (i *FloatObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{Value: i.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *FloatObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{Value: i.Value < o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *FloatObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *FloatObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{Value: i.Value > o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *FloatObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *FloatObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	}

	return &BooleanObject{Value: false}
}
