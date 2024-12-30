package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileLiteral(literal ast.Statement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	switch literal := literal.(type) {
	case *ast.IdentifierLiteral:
		//return c.compileIdentifierLiteral(literal)
	case *ast.IntLiteral:
		return compileIntLiteral(literal, function, vtype, isReturn)
	case *ast.FloatLiteral:
		//return c.compileFloatLiteral(literal)
	case *ast.StringLiteral:
		return compileStringLiteral(c, function, literal)
	case *ast.BooleanLiteral:
		return compileBooleanLiteral(literal, function, vtype, isReturn)
	}

	panic(fmt.Sprintf("'%T' is not a valid literal", literal))
}

func compileIntLiteral(literal *ast.IntLiteral, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	prefix := ""

	if isReturn {
		vtype = function.ReturnType
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
		Type:  qbe.String,
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
		Type:  literal.Type,
		Loc:   literal.Loc,
	}, function, vtype, isReturn)
}
