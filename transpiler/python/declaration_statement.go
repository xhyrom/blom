package python

import (
	"blom/ast"
	"blom/env"
	"fmt"
)

func (t PythonTranspiler) TranspileDeclarationStatement(declaration *ast.VariableDeclarationStatement, environment *env.Environment, indent int) string {
	skip := false

	if environment.Parent != nil && environment.Parent.FindVariable(declaration.Name) != nil {
		if !declaration.Redeclaration && environment.Get(declaration.Name) == nil {
			environment.Set(declaration.Name, &env.IntegerObject{
				Value: int64(indent),
			})

			return fmt.Sprintf("%s_%d = %s\n", declaration.Name, indent, t.TranspileAndFunctionifyStatement(declaration.Value, environment, indent))
		}
	} else {
		skip = true

		environment.Set(declaration.Name, &env.IntegerObject{
			Value: int64(indent),
		})
	}

	if environment.Get(declaration.Name) != nil && !skip {
		return fmt.Sprintf("%s_%d = %s\n", declaration.Name, indent, t.TranspileAndFunctionifyStatement(declaration.Value, environment, indent))
	}

	if environment.Parent.Get(declaration.Name) != nil && !skip {
		id := environment.Parent.Get(declaration.Name).(*env.IntegerObject).Value
		if id > 1 {
			return fmt.Sprintf("%s_%d = %s\n", declaration.Name, id, t.TranspileAndFunctionifyStatement(declaration.Value, environment, indent))
		}
	}

	return fmt.Sprintf("%s = %s\n", declaration.Name, t.TranspileAndFunctionifyStatement(declaration.Value, environment, indent))
}
