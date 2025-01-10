package objects

import "blom/ast"

type StringObject struct {
	value string
}

func NewStringObject(value string) *StringObject {
	return &StringObject{value: value}
}

func (s *StringObject) Type() ast.Type {
	return ast.String
}

func (s *StringObject) Inspect() string {
	return s.value
}

func (s *StringObject) Value() interface{} {
	return s.value
}

func (s *StringObject) SetValue(value interface{}) {
	s.value = value.(string)
}

func (s *StringObject) Add(other Object) Object {
	return nil
}

func (s *StringObject) Subtract(other Object) Object {
	return nil
}

func (s *StringObject) Multiply(other Object) Object {
	return nil
}

func (s *StringObject) Divide(other Object) Object {
	return nil
}

func (s *StringObject) Modulo(other Object) Object {
	return nil
}

func (s *StringObject) BitwiseAnd(other Object) Object {
	return nil
}

func (s *StringObject) BitwiseOr(other Object) Object {
	return nil
}

func (s *StringObject) BitwiseXor(other Object) Object {
	return nil
}

func (s *StringObject) LeftShift(other Object) Object {
	return nil
}

func (s *StringObject) RightShift(other Object) Object {
	return nil
}

func (s *StringObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *StringObject:
		return &BooleanObject{value: s.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (s *StringObject) NotEquals(other Object) Object {
	switch o := other.(type) {
	case *StringObject:
		return &BooleanObject{value: s.value != o.value}
	}

	return &BooleanObject{value: true}
}

func (s *StringObject) LessThan(other Object) Object {
	return nil
}

func (s *StringObject) LessThanOrEqual(other Object) Object {
	return nil
}

func (s *StringObject) GreaterThan(other Object) Object {
	return nil
}

func (s *StringObject) GreaterThanOrEqual(other Object) Object {
	return nil
}
