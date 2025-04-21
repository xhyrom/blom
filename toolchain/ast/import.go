package ast

import "blom/tokens"

type Import struct {
	Path  string
	Alias string
	Loc   tokens.Location
}

func (i Import) Kind() NodeKind {
	return ImportNode
}

func (i Import) Location() tokens.Location {
	return i.Loc
}
