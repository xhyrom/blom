package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileMemberAccess(access *ast.MemberAccess, function *qbe.Function, vtype qbe.Type, isReturn bool) *qbe.TypedValue {
	left := c.compileStatement(access.Left, function, vtype, isReturn)
	return processMemberAccess(c, left, access.Right, function, true)
}

func processMemberAccess(c *Compiler, left *qbe.TypedValue, right ast.Expression, function *qbe.Function, load bool) *qbe.TypedValue {
	for {
		switch right := right.(type) {
		case *ast.IdentifierLiteral:
			field := right.Value
			ty := left.Type

			structName := ty.(qbe.StructBox).Name
			structType := c.Module.GetTypeByName(structName)

			entity := c.Entities[structName]

			offset, fieldType := memberToOffset(c, entity, structType, field)

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

			if !load {
				return &qbe.TypedValue{Type: fieldType, Value: offsetTmp}
			}

			if !fieldType.IsStruct() {
				tmp := c.getTemporaryValue(nil)
				function.LastBlock().AddAssign(
					tmp,
					fieldType,
					qbe.NewLoadInstruction(fieldType, offsetTmp),
				)

				return &qbe.TypedValue{Type: fieldType, Value: tmp}
			}

			return &qbe.TypedValue{Type: fieldType, Value: offsetTmp}
		case *ast.MemberAccess:
			nestedLeft := processMemberAccess(c, left, right.Left, function, load)
			left = nestedLeft
			right = right.Right.(*ast.MemberAccess)

			continue
		default:
			panic(fmt.Sprintf("Unexpected AST node type for field access: %T", right))
		}
	}
}
