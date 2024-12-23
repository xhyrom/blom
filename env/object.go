package env

import "fmt"

type ObjectType int

const (
	Boolean ObjectType = iota
	Integer
	Float
)

type Object interface {
	Type() ObjectType
	Inspect() string
	Add(other Object) Object
	Subtract(other Object) Object
	Multiply(other Object) Object
	Divide(other Object) Object
	Modulo(other Object) Object
	BitwiseAnd(other Object) Object
	BitwiseOr(other Object) Object
	BitwiseXor(other Object) Object
	LeftShift(other Object) Object
	RightShift(other Object) Object
	Equals(other Object) Object
	LessThan(other Object) Object
	LessThanOrEqual(other Object) Object
	GreaterThan(other Object) Object
	GreaterThanOrEqual(other Object) Object
}

type BooleanObject struct {
	Value bool
}

func (b *BooleanObject) Type() ObjectType {
	return Boolean
}

func (b *BooleanObject) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *BooleanObject) Add(other Object) Object {
	return nil
}

func (b *BooleanObject) Subtract(other Object) Object {
	return nil
}

func (b *BooleanObject) Multiply(other Object) Object {
	return nil
}

func (b *BooleanObject) Divide(other Object) Object {
	return nil
}

func (b *BooleanObject) Modulo(other Object) Object {
	return nil
}

func (b *BooleanObject) BitwiseAnd(other Object) Object {
	return nil
}

func (b *BooleanObject) BitwiseOr(other Object) Object {
	return nil
}

func (b *BooleanObject) BitwiseXor(other Object) Object {
	return nil
}

func (b *BooleanObject) LeftShift(other Object) Object {
	return nil
}

func (b *BooleanObject) RightShift(other Object) Object {
	return nil
}

func (b *BooleanObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *BooleanObject:
		return &BooleanObject{Value: b.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (b *BooleanObject) LessThan(other Object) Object {
	return nil
}

func (b *BooleanObject) LessThanOrEqual(other Object) Object {
	return nil
}

func (b *BooleanObject) GreaterThan(other Object) Object {
	return nil
}

func (b *BooleanObject) GreaterThanOrEqual(other Object) Object {
	return nil
}

type IntegerObject struct {
	Value int64
}

func (i *IntegerObject) Type() ObjectType {
	return Integer
}

func (i *IntegerObject) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i IntegerObject) Add(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value + o.Value}
	case *FloatObject:
		return &FloatObject{Value: float64(i.Value) + o.Value}
	}

	return nil
}

func (i *IntegerObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value - o.Value}
	case *FloatObject:
		return &FloatObject{Value: float64(i.Value) - o.Value}
	}

	return nil
}

func (i *IntegerObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value * o.Value}
	case *FloatObject:
		return &FloatObject{Value: float64(i.Value) * o.Value}
	}

	return nil
}

func (i *IntegerObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value / o.Value}
	case *FloatObject:
		return &FloatObject{Value: float64(i.Value) / o.Value}
	}

	return nil
}

func (i *IntegerObject) Modulo(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value % o.Value}
	}

	return nil
}

func (i *IntegerObject) BitwiseAnd(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value & o.Value}
	}

	return nil
}

func (i *IntegerObject) BitwiseOr(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value | o.Value}
	}

	return nil
}

func (i *IntegerObject) BitwiseXor(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value ^ o.Value}
	}

	return nil
}

func (i *IntegerObject) LeftShift(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value << uint(o.Value)}
	}

	return nil
}

func (i *IntegerObject) RightShift(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &IntegerObject{Value: i.Value >> uint(o.Value)}
	}

	return nil
}

func (i *IntegerObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: i.Value == o.Value}
	case *FloatObject:
		return &BooleanObject{Value: float64(i.Value) == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (i *IntegerObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: i.Value < o.Value}
	case *FloatObject:
		return &BooleanObject{Value: float64(i.Value) < o.Value}
	}

	return nil
}

func (i *IntegerObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: i.Value <= o.Value}
	case *FloatObject:
		return &BooleanObject{Value: float64(i.Value) <= o.Value}
	}

	return nil
}

func (i *IntegerObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: i.Value > o.Value}
	case *FloatObject:
		return &BooleanObject{Value: float64(i.Value) > o.Value}
	}

	return nil
}

func (i *IntegerObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: i.Value >= o.Value}
	case *FloatObject:
		return &BooleanObject{Value: float64(i.Value) >= o.Value}
	}

	return nil
}

type FloatObject struct {
	Value float64
}

func (f *FloatObject) Type() ObjectType {
	return Float
}

func (f *FloatObject) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}

func (f *FloatObject) Add(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &FloatObject{Value: f.Value + float64(o.Value)}
	case *FloatObject:
		return &FloatObject{Value: f.Value + o.Value}
	}

	return nil
}

func (f *FloatObject) Subtract(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &FloatObject{Value: f.Value - float64(o.Value)}
	case *FloatObject:
		return &FloatObject{Value: f.Value - o.Value}
	}

	return nil
}

func (f *FloatObject) Multiply(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &FloatObject{Value: f.Value * float64(o.Value)}
	case *FloatObject:
		return &FloatObject{Value: f.Value * o.Value}
	}

	return nil
}

func (f *FloatObject) Divide(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &FloatObject{Value: f.Value / float64(o.Value)}
	case *FloatObject:
		return &FloatObject{Value: f.Value / o.Value}
	}

	return nil
}

func (f *FloatObject) Modulo(other Object) Object {
	return nil
}

func (f *FloatObject) BitwiseAnd(other Object) Object {
	return nil
}

func (f *FloatObject) BitwiseOr(other Object) Object {
	return nil
}

func (f *FloatObject) BitwiseXor(other Object) Object {
	return nil
}

func (f *FloatObject) LeftShift(other Object) Object {
	return nil
}

func (f *FloatObject) RightShift(other Object) Object {
	return nil
}

func (f *FloatObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: f.Value == float64(o.Value)}
	case *FloatObject:
		return &BooleanObject{Value: f.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (f *FloatObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: f.Value < float64(o.Value)}
	case *FloatObject:
		return &BooleanObject{Value: f.Value < o.Value}
	}

	return nil
}

func (f *FloatObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: f.Value <= float64(o.Value)}
	case *FloatObject:
		return &BooleanObject{Value: f.Value <= o.Value}
	}

	return nil
}

func (f *FloatObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: f.Value > float64(o.Value)}
	case *FloatObject:
		return &BooleanObject{Value: f.Value > o.Value}
	}

	return nil
}

func (f *FloatObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *IntegerObject:
		return &BooleanObject{Value: f.Value >= float64(o.Value)}
	case *FloatObject:
		return &BooleanObject{Value: f.Value >= o.Value}
	}

	return nil
}
