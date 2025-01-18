package objects

import (
	"blom/ast"
	"fmt"
)

type PointerObject struct {
	target *Object
}

func NewPointerObject(target Object) *PointerObject {
	return &PointerObject{target: &target}
}

func (p *PointerObject) Type() ast.Type {
	return ast.NewPointerType((*p.target).Type().AsId())
}

func (p *PointerObject) Inspect() string {
	return fmt.Sprintf("&%s", (*p.target).Inspect())
}

func (p *PointerObject) Value() interface{} {
	return *p.target
}

func (p *PointerObject) SetValue(value interface{}) {
	(*p.target).SetValue(value)
}
func (p *PointerObject) Add(other Object) Object        { return nil }
func (p *PointerObject) Subtract(other Object) Object   { return nil }
func (p *PointerObject) Multiply(other Object) Object   { return nil }
func (p *PointerObject) Divide(other Object) Object     { return nil }
func (p *PointerObject) Modulo(other Object) Object     { return nil }
func (p *PointerObject) BitwiseAnd(other Object) Object { return nil }
func (p *PointerObject) BitwiseOr(other Object) Object  { return nil }
func (p *PointerObject) BitwiseXor(other Object) Object { return nil }
func (p *PointerObject) LeftShift(other Object) Object  { return nil }
func (p *PointerObject) RightShift(other Object) Object { return nil }

func (p *PointerObject) Equals(other Object) Object {
	if other, ok := other.(*PointerObject); ok {
		return &BooleanObject{value: p.target == other.target}
	}
	return &BooleanObject{value: false}
}

func (p *PointerObject) NotEquals(other Object) Object {
	if other, ok := other.(*PointerObject); ok {
		return &BooleanObject{value: p.target != other.target}
	}
	return &BooleanObject{value: true}
}

func (p *PointerObject) LessThan(other Object) Object           { return nil }
func (p *PointerObject) LessThanOrEqual(other Object) Object    { return nil }
func (p *PointerObject) GreaterThan(other Object) Object        { return nil }
func (p *PointerObject) GreaterThanOrEqual(other Object) Object { return nil }
