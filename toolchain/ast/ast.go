package ast

import "blom/tokens"

type NodeKind int

const (
	ProgramNode NodeKind = iota
	IdentifierNode
	CharNode
	StringNode
	IntNode
	FloatNode
	BooleanNode

	BinaryExpressionNode
	UnaryExpressionNode
	VariableDeclarationNode
	AssignmentNode
	ReturnNode
	BlockNode
	IfNode
	WhileLoopNode
	AnnotationNode
	FunctionDeclarationNode
	FunctionCallNode
)

type Node interface {
	Kind() NodeKind
	Location() tokens.Location
}

type Program struct {
	Body []Node
}
