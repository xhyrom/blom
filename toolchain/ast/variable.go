package ast

import (
	"blom/tokens"
)

type VariableDeclaration struct {
	Name    IdentifierLiteral
	Type    Type
	Init    Node
	Mutable bool
	Loc     tokens.Location
}

func (d VariableDeclaration) Kind() NodeKind {
	return VariableDeclarationNode
}

func (d VariableDeclaration) Location() tokens.Location {
	return d.Loc
}
