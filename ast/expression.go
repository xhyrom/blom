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
