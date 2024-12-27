package ast

import "blom/tokens"

type NodeKind int

const (
	ProgramNode NodeKind = iota
	IdentifierLiteralNode
	CharLiteralNode
	StringLiteralNode
	IntLiteralNode
	FloatLiteralNode
	BinaryExpressionNode
	UnaryExpressionNode
	DeclarationNode
	ReturnNode
	BlockNode
	IfNode
	ForLoopNode
	WhileLoopNode

	AnnotationNode
	FunctionDeclarationNode
	FunctionCallNode
)

type Statement interface {
	Kind() NodeKind
	Location() tokens.Location
	SetLocation(row uint64, column uint64)
}

type Program BlockStatement
