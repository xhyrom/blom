package ast

import "blom/tokens"

type FunctionArgument struct {
	Name string
	Type int
}

type FunctionDeclaration struct {
	Name        string
	Arguments   []FunctionArgument
	Annotations []Annotation
	ReturnType  int
	Body        *BlockStatement
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
