package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretBuiltinFunctionCall(call *ast.BuiltinFunctionCall) objects.Object {
	switch call.Name {
	case "cast":
		return interpretCastFunctionCall(t, call)
	}

	panic("Unknown builtin function call")
}

func interpretCastFunctionCall(t *Interpreter, call *ast.BuiltinFunctionCall) objects.Object {
	if len(call.Parameters) < 2 {
		panic("Cast function requires at least two parameters")
	}

	firstParam := call.Parameters[0]
	if firstParam.Kind() != ast.IdentifierLiteralNode {
		panic("First parameter of cast function must be an type")
	}

	typeName := firstParam.(*ast.IdentifierLiteral).Value
	astType, err := ast.ParseType(typeName)
	if err != nil {
		panic("Invalid type name")
	}

	stmt := t.interpretStatement(call.Parameters[1])

	return t.convertToType(
		stmt.Type(),
		astType,
		stmt,
	)
}
