package ast

import "blom/tokens"

type DeclarationStatement struct {
	Name  string
	Type  int
	Value Statement
	Loc   tokens.Location
}

func (d DeclarationStatement) Kind() NodeKind {
	return DeclarationNode
}

func (d DeclarationStatement) Location() tokens.Location {
	return d.Loc
}
