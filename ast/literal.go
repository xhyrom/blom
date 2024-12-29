package ast

import (
	"blom/tokens"
)

type IdentifierLiteral struct {
	Value string
	Type  Type
	Loc   tokens.Location
}

func (l IdentifierLiteral) Kind() NodeKind {
	return IdentifierLiteralNode
}

func (l IdentifierLiteral) Location() tokens.Location {
	return l.Loc
}

func (l IdentifierLiteral) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type CharLiteral struct {
	Value rune
	Type  Type
	Loc   tokens.Location
}

func (l CharLiteral) Kind() NodeKind {
	return CharLiteralNode
}

func (l CharLiteral) Location() tokens.Location {
	return l.Loc
}

func (l CharLiteral) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type StringLiteral struct {
	Value string
	Type  Type
	Loc   tokens.Location
}

func (l StringLiteral) Kind() NodeKind {
	return StringLiteralNode
}

func (l StringLiteral) Location() tokens.Location {
	return l.Loc
}

func (l StringLiteral) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type IntLiteral struct {
	Value int64
	Type  Type
	Loc   tokens.Location
}

func (l IntLiteral) Kind() NodeKind {
	return IntLiteralNode
}

func (l IntLiteral) Location() tokens.Location {
	return l.Loc
}

func (l IntLiteral) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type FloatLiteral struct {
	Value float64
	Type  Type
	Loc   tokens.Location
}

func (l FloatLiteral) Kind() NodeKind {
	return FloatLiteralNode
}

func (l FloatLiteral) Location() tokens.Location {
	return l.Loc
}

func (l FloatLiteral) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}

type BooleanLiteral struct {
	Value bool
	Type  Type
	Loc   tokens.Location
}

func (l BooleanLiteral) Kind() NodeKind {
	return BooleanLiteralNode
}

func (l BooleanLiteral) Location() tokens.Location {
	return l.Loc
}

func (l BooleanLiteral) SetLocation(row uint64, column uint64) {
	l.Loc.Row = row
	l.Loc.Column = column
}
