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
	CustomTypes() map[string]ast.Type
	AddCustomType(name string, ty ast.Type)
	ParseStatement() ([]ast.Statement, error)
	ParseExpression() (ast.Expression, error)
	ParsePrimaryExpression() (ast.Expression, error)
}
