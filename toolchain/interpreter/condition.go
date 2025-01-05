package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretCondition(conditionStatement *ast.If, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	t.Scopes.Append()

	// If condition
	condition := t.interpretStatement(conditionStatement.Condition, function, vtype, isReturn)

	if condition.(*objects.BooleanObject).Value().(bool) == true {
		for _, statement := range conditionStatement.Then {
			value := t.interpretStatement(statement, function, vtype, isReturn)
			if _, ok := statement.(*ast.ReturnStatement); ok {
				return value
			}
		}
	} else if conditionStatement.Else != nil {
		for _, statement := range conditionStatement.Else {
			value := t.interpretStatement(statement, function, vtype, isReturn)
			if _, ok := statement.(*ast.ReturnStatement); ok {
				return value
			}
		}
	}

	t.Scopes.Pop()

	return nil
}
