package ast

import "blom/tokens"

type ReturnStatement struct {
	Value Statement
	Loc   tokens.Location
}

func (r ReturnStatement) Kind() NodeKind {
	return ReturnNode
}

func (r ReturnStatement) Location() tokens.Location {
	return r.Loc
}

func (r *ReturnStatement) SetLocation(row uint64, column uint64) {
	r.Loc.Row = row
	r.Loc.Column = column
}
