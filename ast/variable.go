package ast

import (
	"blom/compiler"
	"blom/tokens"
)

type VariableDeclarationStatement struct {
	Name  string
	Type  compiler.Type
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

type AssignmentStatement struct {
	Name  string
	Value Expression
	Loc   tokens.Location
}

func (a AssignmentStatement) Kind() NodeKind {
	return AssignmentNode
}

func (a AssignmentStatement) Location() tokens.Location {
	return a.Loc
}

func (a *AssignmentStatement) SetLocation(row uint64, column uint64) {
	a.Loc.Row = row
	a.Loc.Column = column
}
