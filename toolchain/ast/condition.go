package ast

import "blom/tokens"

type If struct {
	Condition Node
	Then      *Block
	Else      *Block
	Loc       tokens.Location
}

func (i If) Kind() NodeKind {
	return IfNode
}

func (i If) Location() tokens.Location {
	return i.Loc
}
