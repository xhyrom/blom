package expressions

import (
	"blom/ast"
	"blom/env"
)

func InterpretFunctionDeclaration(interpreter Interpreter, environment *env.Environment, declaration *ast.FunctionDeclaration) {
	environment.SetFunction(declaration.Name, declaration)
}
