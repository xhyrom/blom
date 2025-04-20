package ast

import (
	"blom/tokens"
)

type VariableDeclaration struct {
	Id   *IdentifierLiteral
	Type Type
	Init Node
	Loc  tokens.Location
}

func (d VariableDeclaration) Kind() NodeKind {
	return VariableDeclarationNode
}

func (d VariableDeclaration) Location() tokens.Location {
	return d.Loc
}
