package qbe

import (
	"blom/ast"
	"fmt"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileLiteral(literal ast.Node, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
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

func compileIdentifierLiteral(c *Codegen, literal *ast.IdentifierLiteral, function *ir.Function) *ir.TypedValue {
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
		variable.Type.IntoBase(),
		ir.NewLoadInstruction(variable.Type.IntoBase(), address.Value),
	)

	return variable
}

func compileIntLiteral(literal *ast.IntLiteral, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
	prefix := ""

	var t ir.Type = ir.Word
	if vtype != nil {
		t = vtype
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

	return &ir.TypedValue{
		Value: ir.ConstantValue[int64]{
			Value:  literal.Value,
			Prefix: prefix,
		},
		Type: t,
	}
}

func compileFloatLiteral(literal *ast.FloatLiteral, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
	prefix := ""

	var t ir.Type = ir.Single
	if vtype != nil {
		t = vtype
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
	case ir.Double:
		prefix = "d_"
	case ir.Single:
		prefix = "s_"
	}

	return &ir.TypedValue{
		Value: ir.ConstantValue[float64]{
			Value:  literal.Value,
			Prefix: prefix,
		},
		Type: t,
	}
}

func compileCharLiteral(literal *ast.CharLiteral) *ir.TypedValue {
	return &ir.TypedValue{
		Value: ir.ConstantValue[int64]{
			Value: int64(literal.Value),
		},
		Type: ir.Char,
	}
}

func compileStringLiteral(c *Codegen, function *ir.Function, literal *ast.StringLiteral) *ir.TypedValue {
	name := c.assignNameToValueWithPrefix(function.Name)

	c.Module.AddData(ir.Data{
		Linkage: ir.NewLinkage(false),
		Name:    name,
		Items: []ir.TypedDataItem{
			ir.NewTypedDataItem(ir.Byte, ir.NewStringDataItem(literal.Value)),
			ir.NewTypedDataItem(ir.Byte, ir.NewConstantDataItem(0)),
		},
	})

	return &ir.TypedValue{
		Value: ir.NewGlobalValue(name),
		Type:  ir.NewPointer(ir.Char),
	}
}

func boolToInt(value bool) int64 {
	if value {
		return 1
	}

	return 0
}

func compileBooleanLiteral(literal *ast.BooleanLiteral, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
	return compileIntLiteral(&ast.IntLiteral{
		Value: boolToInt(literal.Value),
		Loc:   literal.Loc,
	}, function, vtype, isReturn)
}
