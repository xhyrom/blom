package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileEntity(entity ast.Entity) qbe.TypeDefinition {
	items := make([]qbe.TypedTypeDefinitionItem, len(entity.Fields))

	for i, field := range entity.Fields {
		t := qbe.RemapAstType(field.Type)
		items[i] = qbe.TypedTypeDefinitionItem{
			Count: 1,
			Type:  t,
		}
	}

	return qbe.TypeDefinition{
		Name:  entity.Name,
		Items: items,
	}
}

func (c *Compiler) compileEntityConstruction(construction *ast.EntityConstruction, function *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	entity := c.Entities[construction.Name]
	structType := c.Module.GetTypeByName(construction.Name)

	alloc := c.createVariable(qbe.Long, construction.Name)

	function.LastBlock().AddAssign(alloc, qbe.Long, qbe.NewAlloc8Instruction(qbe.NewConstantValue(4)))

	for name, exp := range construction.Values {
		offset, ty := memberToOffset(entity, structType, name)

		value := c.compileStatement(exp, function, ty, false)

		offsetTmp := c.createVariable(qbe.Long, "offset")
		function.LastBlock().AddAssign(offsetTmp, qbe.Long, qbe.NewAddInstruction(alloc, qbe.NewConstantValue(offset)))

		function.LastBlock().AddInstruction(qbe.NewStoreInstruction(qbe.Word, value.Value, offsetTmp))
	}

	return &qbe.TypedValue{
		Type:  vtype,
		Value: alloc,
	}
}

func memberToOffset(entity *ast.Entity, structType *qbe.TypeDefinition, field string) (int64, qbe.Type) {
	var offset int64
	var fieldType qbe.Type
	for i, item := range entity.Fields {
		if item.Name == field {
			offset = int64(i)
			fieldType = structType.Items[i].Type
			break
		}
	}

	return offset, fieldType
}
