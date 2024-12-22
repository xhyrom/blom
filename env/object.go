package env

import "fmt"

type ObjectType int

const (
	Integer ObjectType = iota
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
