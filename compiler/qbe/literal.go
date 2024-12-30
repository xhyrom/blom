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
		//return c.compileFloatLiteral(literal)
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
	variable := c.getVariable(literal.Value)
	if variable == nil {
		panic("missing variable")
	}

	address := c.getVariable(fmt.Sprintf("%s.addr", literal.Value))
	if address == nil {
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

	if isReturn {
		vtype = &function.ReturnType
	}

	switch *vtype {
	case qbe.Double:
		prefix = "d_"
	case qbe.Single:
		prefix = "s_"
	}

	return &qbe.TypedValue{
		Value: qbe.ConstantValue{
			Value:  literal.Value,
			Prefix: prefix,
		},
		Type: *vtype,
	}
}

func compileCharLiteral(literal *ast.CharLiteral) *qbe.TypedValue {
	return &qbe.TypedValue{
		Value: qbe.ConstantValue{
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
