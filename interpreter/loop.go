package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretLoop(loopStatement *ast.WhileLoopStatement, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	t.Scopes.Append()

	// Loop condition
	condition := t.interpretStatement(loopStatement.Condition, function, vtype, isReturn)

	for condition.(*objects.BooleanObject).Value().(bool) {
		// Loop body
		for _, statement := range loopStatement.Body {
			t.interpretStatement(statement, function, vtype, isReturn)
		}

		condition = t.interpretStatement(loopStatement.Condition, function, vtype, isReturn)
	}

	t.Scopes.Pop()

	return nil
}
