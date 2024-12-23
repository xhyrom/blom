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
