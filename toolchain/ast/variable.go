package ast

import (
	"blom/tokens"
)

type VariableDeclaration struct {
	Name        string
	Type        Type
	Value       Expression
	Annotations []Annotation
	Loc         tokens.Location
}

func (d VariableDeclaration) Kind() NodeKind {
	return VariableDeclarationNode
}

func (d VariableDeclaration) Category() Category {
	return StatementCategory
}

func (d VariableDeclaration) Location() tokens.Location {
	return d.Loc
}

func (d *VariableDeclaration) SetLocation(row uint64, column uint64) {
	d.Loc.Row = row
	d.Loc.Column = column
}

type Assignment struct {
	Left  Expression
	Right Expression
	Loc   tokens.Location
}

func (a Assignment) Kind() NodeKind {
	return AssignmentNode
}

func (a Assignment) Category() Category {
	return StatementCategory
}

func (a Assignment) Location() tokens.Location {
	return a.Loc
}

func (a *Assignment) SetLocation(row uint64, column uint64) {
	a.Loc.Row = row
	a.Loc.Column = column
}
