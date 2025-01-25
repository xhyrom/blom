package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretBuiltinFunctionCall(call *ast.BuiltinFunctionCall, function *ast.FunctionDeclaration, vtype *ast.Type) objects.Object {
	switch call.Name {
	case "cast":
		return interpretCastFunctionCall(t, call, function, vtype)
	}

	panic("Unknown builtin function call")
}

func interpretCastFunctionCall(t *Interpreter, call *ast.BuiltinFunctionCall, function *ast.FunctionDeclaration, vtype *ast.Type) objects.Object {
	if len(call.Parameters) < 2 {
		panic("Cast function requires at least two parameters")
	}

	firstParam := call.Parameters[0]
	if firstParam.Kind() != ast.IdentifierLiteralNode {
		panic("First parameter of cast function must be an type")
	}

	typeName := firstParam.(*ast.IdentifierLiteral).Value
	astType, err := ast.ParseType(typeName, map[string]ast.Type{})
	if err != nil {
		panic("Invalid type name")
	}

	stmt := t.interpretStatement(call.Parameters[1], function, vtype, false)

	return t.convertToType(
		stmt.Type(),
		astType,
		stmt,
	)
}
