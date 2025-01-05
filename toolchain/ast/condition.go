package ast

import "blom/tokens"

type If struct {
	Condition Expression
	Then      []Statement
	Else      []Statement
	Loc       tokens.Location
}

func (i If) Kind() NodeKind {
	return IfNode
}

func (i If) Location() tokens.Location {
	return i.Loc
}

func (i *If) SetLocation(row uint64, column uint64) {
	i.Loc.Row = row
	i.Loc.Column = column
}

func (i *If) HasElse() bool {
	if i.Else != nil && len(i.Else) > 0 {
		return true
	}

	return false
}
