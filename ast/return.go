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
