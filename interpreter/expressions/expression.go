package expressions

import (
	"blom/ast"
	"blom/env"
)

type Interpreter interface {
	InterpretBlock(block ast.BlockStatement, environment *env.Environment) env.Object
	InterpretStatement(statement ast.Statement, environment *env.Environment) env.Object
	InterpretExpression(expression ast.Expression, environment *env.Environment) env.Object
}
