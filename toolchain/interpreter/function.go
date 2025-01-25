package interpreter

import (
	"blom/ast"
	"blom/interpreter/native"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretFunctionCall(call *ast.FunctionCall, currentFunction *ast.FunctionDeclaration, vtype *ast.Type) objects.Object {
	function := t.Functions[call.Name]
	if function == nil {
		// lambda
		lambda, exists := t.Scopes.GetValue(call.Name)
		if !exists {
			panic("missing function")
		}

		if lambda.Type().IsPointer() {
			function = lambda.(*objects.PointerObject).Value().(*objects.LambdaObject).AsFunction()
		} else {
			function = lambda.(*objects.LambdaObject).AsFunction()
		}
	}

	if function.IsNative() {
		objects := make([]objects.Object, len(call.Parameters))
		for i, parameter := range call.Parameters {
			var argType *ast.Type
			if i < len(function.Arguments) {
				argType = &function.Arguments[i].Type
			} else {
				argType = vtype
			}

			objects[i] = t.interpretStatement(parameter, currentFunction, argType, false)
		}

		native.Printf(objects)
		return nil
	}

	t.Scopes.Append()
	startScopeIndex := t.Scopes.Length() - 1

	for i, parameter := range call.Parameters {
		t.Scopes.Set(function.Arguments[i].Name, t.interpretStatement(parameter, currentFunction, &function.Arguments[i].Type, false))
	}

	for _, statement := range function.Body {
		value := t.interpretStatement(statement, function, &function.ReturnType, false)
		switch statement.(type) {
		case *ast.ReturnStatement:
			t.Scopes.PopUntil(startScopeIndex)
			return value
		case *ast.If:
			if value != nil {
				t.Scopes.PopUntil(startScopeIndex)
				return value
			}
		}
	}

	t.Scopes.PopUntil(startScopeIndex)

	return nil
}
