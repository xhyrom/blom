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
	BooleanLiteralNode
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
	CompileTimeFunctionCallNode
	TypeDefinitionNode

	EntityNode
	EntityConstructionNode
	MemberAccessNode
)

type Category int

const (
	ExpressionCategory Category = iota
	StatementCategory
)

type Statement interface {
	Kind() NodeKind
	Category() Category
	Location() tokens.Location
	SetLocation(row uint64, column uint64)
}

type Program Block
