package ast

import "blom/tokens"

type Return struct {
	Value Node
	Loc   tokens.Location
}

func (r Return) Kind() NodeKind {
	return ReturnNode
}

func (r Return) Location() tokens.Location {
	return r.Loc
}
