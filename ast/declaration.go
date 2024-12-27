package ast

import "blom/tokens"

type DeclarationStatement struct {
	Name          string
	Type          int
	Value         Statement
	Redeclaration bool
	Loc           tokens.Location
}

func (d DeclarationStatement) Kind() NodeKind {
	return DeclarationNode
}

func (d DeclarationStatement) Location() tokens.Location {
	return d.Loc
}

func (d *DeclarationStatement) SetLocation(row uint64, column uint64) {
	d.Loc.Row = row
	d.Loc.Column = column
}
