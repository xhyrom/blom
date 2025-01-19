package ast

import "blom/tokens"

type TypeDefinition struct {
	Name string
	Type Type
	Loc  tokens.Location
}

func (t TypeDefinition) Kind() NodeKind {
	return TypeDefinitionNode
}

func (t TypeDefinition) Location() tokens.Location {
	return t.Loc
}

func (t *TypeDefinition) SetLocation(row uint64, column uint64) {
	t.Loc.Row = row
	t.Loc.Column = column
}
