package ast

import (
	"blom/tokens"
	"strings"
)

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

func (p Path) Last() IdentifierLiteral {
	if len(p.Segments) == 0 {
		return IdentifierLiteral{}
	}

	return p.Segments[len(p.Segments)-1]
}

func (p Path) Dotify() string {
	segments := make([]string, len(p.Segments))
	for i, segment := range p.Segments {
		segments[i] = segment.Value
	}

	return strings.Join(segments, ".")
}
