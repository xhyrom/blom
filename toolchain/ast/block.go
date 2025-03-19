package ast

import "blom/tokens"

type Block struct {
	Body []Statement
	Cat  Category
	Loc  tokens.Location
}

func (b Block) Kind() NodeKind {
	return BlockNode
}

func (b Block) Category() Category {
	return b.Cat
}

func (b Block) Location() tokens.Location {
	return b.Loc
}

func (b *Block) SetLocation(row uint64, column uint64) {
	b.Loc.Row = row
	b.Loc.Column = column
}
