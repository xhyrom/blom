package expressions

import (
	"blom/ast"
	"blom/tokens"
)

type Parser interface {
	Source() string
	IsEof() bool
	Current() tokens.Token
	Next() tokens.Token
	Consume() tokens.Token
	ParseStatement() ([]ast.Statement, error)
	ParseExpression() (ast.Expression, error)
	ParsePrimaryExpression() (ast.Expression, error)
}
