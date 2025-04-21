package ast

import "blom/tokens"

type Call interface {
	Node
	isCall()
}

type FunctionCall struct {
	Path Path
	Args []Node
	Loc  tokens.Location
}

func (f FunctionCall) Kind() NodeKind {
	return FunctionCallNode
}

func (f FunctionCall) Location() tokens.Location {
	return f.Loc
}

func (f FunctionCall) isCall() {}

type MethodCall struct {
	Callee Node
	Args   []Node
	Loc    tokens.Location
}

func (m MethodCall) Kind() NodeKind {
	return MethodCallNode
}

func (m MethodCall) Location() tokens.Location {
	return m.Loc
}

func (m MethodCall) isCall() {}

type InfixCall struct {
	*FunctionCall
	*MethodCall
	Loc tokens.Location
}

func (i InfixCall) Kind() NodeKind {
	return InfixCallNode
}

func (i InfixCall) Location() tokens.Location {
	return i.Loc
}

func (i InfixCall) isCall() {}

// Explicit interface implementation
var _ Call = (*FunctionCall)(nil)
var _ Call = (*MethodCall)(nil)
var _ Call = (*InfixCall)(nil)
