package ast

import "blom/tokens"

type IdentifierLiteralStatement struct {
	Value string
	Loc   tokens.Location
}

func (l IdentifierLiteralStatement) Kind() NodeKind {
	return IdentifierLiteralNode
}

func (l IdentifierLiteralStatement) Location() tokens.Location {
	return l.Loc
}

type IntLiteralStatement struct {
	Value int64
	Loc   tokens.Location
}

func (l IntLiteralStatement) Kind() NodeKind {
	return IntLiteralNode
}

func (l IntLiteralStatement) Location() tokens.Location {
	return l.Loc
}

type FloatLiteralStatement struct {
	Value float64
	Loc   tokens.Location
}

func (l FloatLiteralStatement) Kind() NodeKind {
	return FloatLiteralNode
}

func (l FloatLiteralStatement) Location() tokens.Location {
	return l.Loc
}
