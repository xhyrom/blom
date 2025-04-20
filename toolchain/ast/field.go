package ast

import "blom/tokens"

type Field struct {
	Base Node
	Member Node
	Loc tokens.Location
}

func (f Field) Kind() NodeKind {
	return FieldNode
}

func (f Field) Location() tokens.Location {
	return f.Loc
}
