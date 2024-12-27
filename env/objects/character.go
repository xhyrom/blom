package objects

import (
	"blom/compiler"
	"fmt"
)

type CharacterObject struct {
	Value rune
}

func (c *CharacterObject) Type() compiler.Type {
	return compiler.Char
}

func (c *CharacterObject) Inspect() string {
	return fmt.Sprintf("%c", c.Value)
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
		return &BooleanObject{Value: c.Value == o.Value}
	}

	return &BooleanObject{Value: false}
}

func (c *CharacterObject) LessThan(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{Value: c.Value < o.Value}
	}

	return nil
}

func (c *CharacterObject) LessThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{Value: c.Value <= o.Value}
	}

	return nil
}

func (c *CharacterObject) GreaterThan(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{Value: c.Value > o.Value}
	}

	return nil
}

func (c *CharacterObject) GreaterThanOrEqual(other Object) Object {
	switch o := other.(type) {
	case *CharacterObject:
		return &BooleanObject{Value: c.Value >= o.Value}
	}

	return nil
}
