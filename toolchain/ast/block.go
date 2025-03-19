package ast

import "blom/tokens"

type BlockStatement struct {
	Body []Statement
	Loc  tokens.Location
}

func (b BlockStatement) Kind() NodeKind {
	return BlockNode
}

func (b BlockStatement) Location() tokens.Location {
	return b.Loc
}

func (b *BlockStatement) SetLocation(row uint64, column uint64) {
	b.Loc.Row = row
	b.Loc.Column = column
}

type BlockExpression struct {
	Body []Statement
	Loc  tokens.Location
}

func (b BlockExpression) Kind() NodeKind {
	return BlockNode
}

func (b BlockExpression) Location() tokens.Location {
	return b.Loc
}

func (b *BlockExpression) SetLocation(row uint64, column uint64) {
	b.Loc.Row = row
	b.Loc.Column = column
}
