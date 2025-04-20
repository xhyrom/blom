package ast

import (
	"blom/tokens"
)

type IdentifierLiteral struct {
	Value string
	Loc   tokens.Location
}

func (l IdentifierLiteral) Kind() NodeKind {
	return IdentifierNode
}

func (l IdentifierLiteral) Location() tokens.Location {
	return l.Loc
}

type CharLiteral struct {
	Value rune
	Loc   tokens.Location
}

func (l CharLiteral) Kind() NodeKind {
	return CharNode
}

func (l CharLiteral) Location() tokens.Location {
	return l.Loc
}

type StringLiteral struct {
	Value string
	Loc   tokens.Location
}

func (l StringLiteral) Kind() NodeKind {
	return StringNode
}

func (l StringLiteral) Location() tokens.Location {
	return l.Loc
}

type IntLiteral struct {
	Value int64
	Loc   tokens.Location
}

func (l IntLiteral) Kind() NodeKind {
	return IntNode
}

func (l IntLiteral) Location() tokens.Location {
	return l.Loc
}

type FloatLiteral struct {
	Value float64
	Loc   tokens.Location
}

func (l FloatLiteral) Kind() NodeKind {
	return FloatNode
}

func (l FloatLiteral) Location() tokens.Location {
	return l.Loc
}

type BooleanLiteral struct {
	Value bool
	Loc   tokens.Location
}

func (l BooleanLiteral) Kind() NodeKind {
	return BooleanNode
}

func (l BooleanLiteral) Location() tokens.Location {
	return l.Loc
}
