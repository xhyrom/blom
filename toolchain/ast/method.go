package ast

import "blom/tokens"

type MethodCall struct {
	Receiver Node
	Path     Path
	Args     []Node
	Loc      tokens.Location
}

func (m MethodCall) Kind() NodeKind {
	return MethodCallNode
}

func (m MethodCall) Location() tokens.Location {
	return m.Loc
}
