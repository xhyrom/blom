package ast

import "blom/tokens"

type IfStatement struct {
	Condition Expression
	Then      Statement
	Else      Statement
	Loc       tokens.Location
}

func (i IfStatement) Kind() NodeKind {
	return IfNode
}

func (i IfStatement) Location() tokens.Location {
	return i.Loc
}
