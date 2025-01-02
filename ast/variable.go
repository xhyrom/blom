package ast

import (
	"blom/tokens"
)

type VariableDeclarationStatement struct {
	Name  string
	Type  Type
	Value Expression
	Loc   tokens.Location
}

func (d VariableDeclarationStatement) Kind() NodeKind {
	return VariableDeclarationNode
}

func (d VariableDeclarationStatement) Location() tokens.Location {
	return d.Loc
}

func (d *VariableDeclarationStatement) SetLocation(row uint64, column uint64) {
	d.Loc.Row = row
	d.Loc.Column = column
}

type Assignment struct {
	Name  string
	Value Expression
	Loc   tokens.Location
}

func (a Assignment) Kind() NodeKind {
	return AssignmentNode
}

func (a Assignment) Location() tokens.Location {
	return a.Loc
}

func (a *Assignment) SetLocation(row uint64, column uint64) {
	a.Loc.Row = row
	a.Loc.Column = column
}
