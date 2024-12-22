package expressions

import (
	"blom/ast"
	"blom/tokens"
)

type Parser interface {
	IsEof() bool
	Current() tokens.Token
	Consume() tokens.Token
	ParseStatement() ast.Statement
	ParseExpression() ast.Expression
}
