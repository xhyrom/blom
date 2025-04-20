package ast

import "blom/tokens"

type BinaryExpression struct {
	Left     Node
	Right    Node
	Operator tokens.TokenKind
	Loc      tokens.Location
}

func (b BinaryExpression) Kind() NodeKind {
	return BinaryExpressionNode
}

func (b BinaryExpression) Location() tokens.Location {
	return b.Loc
}

type UnaryExpression struct {
	Operand  Node
	Operator tokens.TokenKind
	Loc      tokens.Location
}

func (u UnaryExpression) Kind() NodeKind {
	return UnaryExpressionNode
}

func (u UnaryExpression) Location() tokens.Location {
	return u.Loc
}

type Assignment struct {
	Left  Node
	Right Node
	Loc   tokens.Location
}

func (a Assignment) Kind() NodeKind {
	return AssignmentNode
}

func (a Assignment) Location() tokens.Location {
	return a.Loc
}
