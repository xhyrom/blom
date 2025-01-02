package ast

import (
	"blom/tokens"
	"strings"
)

type FunctionArgument struct {
	Name string
	Type Type
}

type FunctionDeclaration struct {
	Name        string
	Arguments   []FunctionArgument
	Annotations []Annotation
	ReturnType  Type
	Body        []Statement
	Variadic    bool
	Loc         tokens.Location
}

func (f FunctionDeclaration) Kind() NodeKind {
	return FunctionDeclarationNode
}

func (f FunctionDeclaration) Location() tokens.Location {
	return f.Loc
}

func (f *FunctionDeclaration) SetLocation(row uint64, column uint64) {
	f.Loc.Row = row
	f.Loc.Column = column
}

func (f *FunctionDeclaration) HasAnnotation(typ AnnotationType) bool {
	for _, annotation := range f.Annotations {
		if annotation.Type == typ {
			return true
		}
	}

	return false
}

func (f *FunctionDeclaration) IsNative() bool {
	return f.HasAnnotation(Native)
}

func (f FunctionDeclaration) PrettyName() string {
	if dotIndex := strings.Index(f.Name, "."); dotIndex != -1 {
		return f.Name[:dotIndex]
	}

	return f.Name
}

type FunctionCall struct {
	Name       string
	Parameters []Expression
	Loc        tokens.Location
}

func (f FunctionCall) Kind() NodeKind {
	return FunctionCallNode
}

func (f FunctionCall) Location() tokens.Location {
	return f.Loc
}

func (f *FunctionCall) SetLocation(row uint64, column uint64) {
	f.Loc.Row = row
	f.Loc.Column = column
}

func (f FunctionCall) PrettyName() string {
	if dotIndex := strings.Index(f.Name, "."); dotIndex != -1 {
		return f.Name[:dotIndex]
	}

	return f.Name
}

type BuiltinFunctionCall struct {
	Name       string
	Parameters []Expression
	Loc        tokens.Location
}

func (c BuiltinFunctionCall) Kind() NodeKind {
	return CompileTimeFunctionCallNode
}

func (c BuiltinFunctionCall) Location() tokens.Location {
	return c.Loc
}

func (c *BuiltinFunctionCall) SetLocation(row uint64, column uint64) {
	c.Loc.Row = row
	c.Loc.Column = column
}
