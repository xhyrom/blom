package ast

import "blom/tokens"

type Return struct {
	Value Expression
	Loc   tokens.Location
}

func (r Return) Kind() NodeKind {
	return ReturnNode
}

func (r Return) Category() Category {
	return StatementCategory
}

func (r Return) Location() tokens.Location {
	return r.Loc
}

func (r *Return) SetLocation(row uint64, column uint64) {
	r.Loc.Row = row
	r.Loc.Column = column
}
