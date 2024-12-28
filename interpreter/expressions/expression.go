package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

type Interpreter interface {
	Source() string
	InterpretBlock(block *ast.BlockStatement, environment *env.Environment[objects.Object]) objects.Object
	InterpretStatement(statement ast.Statement, environment *env.Environment[objects.Object]) objects.Object
}
