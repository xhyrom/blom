package objects

import (
	"blom/ast"
	"fmt"
)

type CharacterObject struct {
	value rune
}

func NewCharacterObject(value rune) *CharacterObject {
	return &CharacterObject{value: value}
}

func (c *CharacterObject) Type() ast.Type {
	return ast.Char
}

func (c *CharacterObject) Inspect() string {
	return fmt.Sprintf("%c", c.value)
}

func (c *CharacterObject) Value() interface{} {
	return c.value
}

func (c *CharacterObject) SetValue(value interface{}) {
	c.value = value.(rune)
}

func (c *CharacterObject) Add(other Object) Object {
	return nil
}

func (c *CharacterObject) Subtract(other Object) Object {
	return nil
}

func (c *CharacterObject) Multiply(other Object) Object {
	return nil
}

func (c *CharacterObject) Divide(other Object) Object {
	return nil
}

func (c *CharacterObject) Modulo(other Object) Object {
	return nil
}

func (c *CharacterObject) BitwiseAnd(other Object) Object {
	return nil
}

func (c *CharacterObject) BitwiseOr(other Object) Object {
	return nil
}

func (c *CharacterObject) BitwiseXor(other Object) Object {
	return nil
}

func (c *CharacterObject) LeftShift(other Object) Object {
	return nil
}

func (c *CharacterObject) RightShift(other Object) Object {
	return nil
}

func (c *CharacterObject) Equals(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{value: c.value == o.value}
	}

	return &BooleanObject{value: false}
}

func (c *CharacterObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{value: c.value < o.value}
	}

	return nil
}

func (c *CharacterObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{value: c.value <= o.value}
	}

	return nil
}

func (c *CharacterObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{value: c.value > o.value}
	}

	return nil
}

func (c *CharacterObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{value: c.value >= o.value}
	}

	return nil
}
