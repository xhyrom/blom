package ast

import "blom/tokens"

type Path struct {
	Segments []IdentifierLiteral
	Loc      tokens.Location
}

func (p Path) Kind() NodeKind {
	return PathNode
}

func (p Path) Location() tokens.Location {
	return p.Loc
}
