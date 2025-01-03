package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretFunctionCall(call *ast.FunctionCall) objects.Object {
	function := t.Functions[call.Name]

	for i, parameter := range call.Parameters {
		t.Scopes.Set(function.Arguments[i].Name, t.interpretStatement(parameter))
	}

	for _, statement := range function.Body {
		t.interpretStatement(statement)
	}

	return nil
}
