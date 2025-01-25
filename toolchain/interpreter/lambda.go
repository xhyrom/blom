package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretLambdaDeclaration(statement *ast.LambdaDeclaration, function *ast.FunctionDeclaration, vtype *ast.Type) objects.Object {
	lambda := objects.NewLambdaObject(*statement)

	return lambda
}
