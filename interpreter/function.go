package interpreter

import (
	"blom/ast"
	"blom/interpreter/native"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretFunctionCall(call *ast.FunctionCall) objects.Object {
	function := t.Functions[call.Name]

	if function.IsNative() {
		objects := make([]objects.Object, len(call.Parameters))
		for i, parameter := range call.Parameters {
			objects[i] = t.interpretStatement(parameter)
		}

		native.Printf(objects)
		return nil
	}

	for i, parameter := range call.Parameters {
		t.Scopes.Set(function.Arguments[i].Name, t.interpretStatement(parameter))
	}

	for _, statement := range function.Body {
		t.interpretStatement(statement)
	}

	return nil
}
