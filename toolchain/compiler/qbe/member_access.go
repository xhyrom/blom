package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileMemberAccess(access *ast.MemberAccess, function *qbe.Function, vtype qbe.Type, isReturn bool) *qbe.TypedValue {
	left := c.compileStatement(access.Left, function, vtype, isReturn)
	value, _ := processMemberAccess(c, left, access.Right, function)

	return value
}

func processMemberAccess(c *Compiler, left *qbe.TypedValue, right ast.Expression, function *qbe.Function) (*qbe.TypedValue, *qbe.TypedValue) {
	for {
		switch right := right.(type) {
		case *ast.IdentifierLiteral:
			field := right.Value
			ty := left.Type

			// Handle pointer automatic dereferencing
			if ty.IsPointer() {
				ty = ty.(qbe.PointerBox).Inner
			}

			structName := ty.(qbe.StructBox).Name
			structType := c.Module.GetTypeByName(structName)

			// Find field offset in struct
			entity := c.Entities[structName]

			offset, fieldType := memberToOffset(entity, structType, field)

			if fieldType == nil {
				panic(fmt.Sprintf("Unknown field '%s' in struct '%s'", field, structName))
			}

			offsetName := "offset"
			offsetTmp := c.getTemporaryValue(&offsetName)
			function.LastBlock().AddAssign(
				offsetTmp,
				qbe.Long,
				qbe.NewAddInstruction(left.Value, qbe.NewConstantValue(offset)),
			)

			if !fieldType.IsStruct() {
				tmp := c.getTemporaryValue(nil)
				function.LastBlock().AddAssign(
					tmp,
					fieldType,
					qbe.NewLoadInstruction(fieldType, offsetTmp),
				)

				return &qbe.TypedValue{Type: fieldType, Value: tmp}, &qbe.TypedValue{Type: qbe.Long, Value: offsetTmp}
			}

			return &qbe.TypedValue{Type: fieldType, Value: offsetTmp}, &qbe.TypedValue{Type: qbe.Long, Value: offsetTmp}
		case *ast.MemberAccess:
			nestedLeft, _ := processMemberAccess(c, left, right.Left, function)
			left = nestedLeft
			right = right.Right.(*ast.MemberAccess)

			continue
		default:
			panic(fmt.Sprintf("Unexpected AST node type for field access: %T", right))
		}
	}
}
