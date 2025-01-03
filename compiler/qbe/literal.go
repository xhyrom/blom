package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileLiteral(literal ast.Statement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	switch literal := literal.(type) {
	case *ast.IdentifierLiteral:
		return compileIdentifierLiteral(c, literal, function)
	case *ast.IntLiteral:
		return compileIntLiteral(literal, function, vtype, isReturn)
	case *ast.FloatLiteral:
		return compileFloatLiteral(literal, function, vtype, isReturn)
	case *ast.CharLiteral:
		return compileCharLiteral(literal)
	case *ast.StringLiteral:
		return compileStringLiteral(c, function, literal)
	case *ast.BooleanLiteral:
		return compileBooleanLiteral(literal, function, vtype, isReturn)
	}

	panic(fmt.Sprintf("'%T' is not a valid literal", literal))
}

func compileIdentifierLiteral(c *Compiler, literal *ast.IdentifierLiteral, function *qbe.Function) *qbe.TypedValue {
	variable, exists := c.Scopes.GetValue(literal.Value)
	if !exists {
		panic("missing variable")
	}

	address, exists := c.Scopes.GetValue(fmt.Sprintf("%s.addr", literal.Value))
	if !exists {
		return variable
	}

	function.LastBlock().AddAssign(
		variable.Value,
		variable.Type,
		qbe.NewLoadInstruction(variable.Type, address.Value),
	)

	return variable
}

func compileIntLiteral(literal *ast.IntLiteral, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	prefix := ""

	var t qbe.Type = qbe.Word
	if vtype != nil {
		t = *vtype
	}

	// compile time casting (int to float)
	if t.IsFloatingPoint() {
		return compileFloatLiteral(&ast.FloatLiteral{
			Value: float64(literal.Value),
			Loc:   literal.Loc,
		}, function, vtype, isReturn)
	}

	if isReturn {
		t = function.ReturnType
	}

	return &qbe.TypedValue{
		Value: qbe.ConstantValue[int64]{
			Value:  literal.Value,
			Prefix: prefix,
		},
		Type: t,
	}
}

func compileFloatLiteral(literal *ast.FloatLiteral, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	prefix := ""

	var t qbe.Type = qbe.Single
	if vtype != nil {
		t = *vtype
	}

	// compile time casting (float to int)
	if t.IsInteger() {
		return compileIntLiteral(&ast.IntLiteral{
			Value: int64(literal.Value),
			Loc:   literal.Loc,
		}, function, vtype, isReturn)
	}

	if isReturn {
		t = function.ReturnType
	}

	switch t {
	case qbe.Double:
		prefix = "d_"
	case qbe.Single:
		prefix = "s_"
	}

	return &qbe.TypedValue{
		Value: qbe.ConstantValue[float64]{
			Value:  literal.Value,
			Prefix: prefix,
		},
		Type: t,
	}
}

func compileCharLiteral(literal *ast.CharLiteral) *qbe.TypedValue {
	return &qbe.TypedValue{
		Value: qbe.ConstantValue[int64]{
			Value: int64(literal.Value),
		},
		Type: qbe.Char,
	}
}

func compileStringLiteral(c *Compiler, function *qbe.Function, literal *ast.StringLiteral) *qbe.TypedValue {
	name := c.assignNameToValueWithPrefix(function.Name)

	c.Module.AddData(qbe.Data{
		Linkage: qbe.NewLinkage(false),
		Name:    name,
		Items: []qbe.TypedDataItem{
			qbe.NewTypedDataItem(qbe.Byte, qbe.NewStringDataItem(literal.Value)),
			qbe.NewTypedDataItem(qbe.Byte, qbe.NewConstantDataItem(0)),
		},
	})

	return &qbe.TypedValue{
		Value: qbe.NewGlobalValue(name),
		Type:  qbe.NewPointer(qbe.Char),
	}
}

func boolToInt(value bool) int64 {
	if value {
		return 1
	}

	return 0
}

func compileBooleanLiteral(literal *ast.BooleanLiteral, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	return compileIntLiteral(&ast.IntLiteral{
		Value: boolToInt(literal.Value),
		Loc:   literal.Loc,
	}, function, vtype, isReturn)
}
