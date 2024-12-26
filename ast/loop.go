package ast

import "blom/tokens"

type ForLoopStatement struct {
	Declaration *DeclarationStatement
	Condition   BinaryExpression
	Step        DeclarationStatement
	Body        BlockStatement
	Loc         tokens.Location
}

func (f ForLoopStatement) Kind() NodeKind {
	return ForLoopNode
}

func (f ForLoopStatement) Location() tokens.Location {
	return f.Loc
}

type WhileLoopStatement struct {
	Condition Expression
	Body      BlockStatement
	Loc       tokens.Location
}

func (w WhileLoopStatement) Kind() NodeKind {
	return WhileLoopNode
}

func (w WhileLoopStatement) Location() tokens.Location {
	return w.Loc
}
