package ast

import "blom/tokens"

type Block struct {
	Body []Node
	Loc  tokens.Location
}

func (b Block) Kind() NodeKind {
	return BlockNode
}

func (b Block) Location() tokens.Location {
	return b.Loc
}
