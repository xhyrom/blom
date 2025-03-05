package ast

import "blom/tokens"

type Expression Statement

type BinaryExpression struct {
	Left     Expression
	Right    Expression
	Operator tokens.TokenKind
	Loc      tokens.Location
}

func (b BinaryExpression) Kind() NodeKind {
	return BinaryExpressionNode
}

func (b BinaryExpression) Location() tokens.Location {
	return b.Loc
}

func (b *BinaryExpression) SetLocation(row uint64, column uint64) {
	b.Loc.Row = row
	b.Loc.Column = column
}

type UnaryExpression struct {
	Operand  Expression
	Operator tokens.TokenKind
	Loc      tokens.Location
}

func (u UnaryExpression) Kind() NodeKind {
	return UnaryExpressionNode
}

func (u UnaryExpression) Location() tokens.Location {
	return u.Loc
}

func (u *UnaryExpression) SetLocation(row uint64, column uint64) {
	u.Loc.Row = row
	u.Loc.Column = column
}
