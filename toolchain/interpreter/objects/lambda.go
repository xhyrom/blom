package objects

import (
	"blom/ast"
	"fmt"
)

type LambdaObject struct {
	declaration *ast.LambdaDeclaration
}

func NewLambdaObject(declaration ast.LambdaDeclaration) *LambdaObject {
	return &LambdaObject{declaration: &declaration}
}

func (p *LambdaObject) AsFunction() *ast.FunctionDeclaration {
	return &ast.FunctionDeclaration{
		Arguments:   p.declaration.Arguments,
		Annotations: []ast.Annotation{},
		ReturnType:  p.declaration.ReturnType,
		Body:        p.declaration.Body,
	}
}

func (p *LambdaObject) Type() ast.Type {
	return ast.Function
}

func (p *LambdaObject) Inspect() string {
	return fmt.Sprintf("lambda")
}

func (p *LambdaObject) Value() interface{} {
	return *p.declaration
}

func (p *LambdaObject) SetValue(value interface{}) {
	*p.declaration = value.(ast.LambdaDeclaration)
}
func (p *LambdaObject) Add(other Object) Object        { return nil }
func (p *LambdaObject) Subtract(other Object) Object   { return nil }
func (p *LambdaObject) Multiply(other Object) Object   { return nil }
func (p *LambdaObject) Divide(other Object) Object     { return nil }
func (p *LambdaObject) Modulo(other Object) Object     { return nil }
func (p *LambdaObject) BitwiseAnd(other Object) Object { return nil }
func (p *LambdaObject) BitwiseOr(other Object) Object  { return nil }
func (p *LambdaObject) BitwiseXor(other Object) Object { return nil }
func (p *LambdaObject) LeftShift(other Object) Object  { return nil }
func (p *LambdaObject) RightShift(other Object) Object { return nil }

func (p *LambdaObject) Equals(other Object) Object {
	if other, ok := other.(*LambdaObject); ok {
		return &BooleanObject{value: p.declaration == other.declaration}
	}
	return &BooleanObject{value: false}
}

func (p *LambdaObject) NotEquals(other Object) Object {
	if other, ok := other.(*LambdaObject); ok {
		return &BooleanObject{value: p.declaration != other.declaration}
	}
	return &BooleanObject{value: true}
}

func (p *LambdaObject) LessThan(other Object) Object           { return nil }
func (p *LambdaObject) LessThanOrEqual(other Object) Object    { return nil }
func (p *LambdaObject) GreaterThan(other Object) Object        { return nil }
func (p *LambdaObject) GreaterThanOrEqual(other Object) Object { return nil }
