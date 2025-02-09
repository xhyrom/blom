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

	alloc := c.getTemporaryValue(&construction.Name)

	function.LastBlock().AddAssign(alloc, qbe.Long, qbe.NewAlloc8Instruction(qbe.NewConstantValue(structType.Size(c.Module))))

	for name, exp := range construction.Values {
		setEntityField(c, function, entity, alloc, name, exp)
	}

	for _, field := range entity.Fields {
		if field.Value == nil {
			continue
		}

		if _, ok := construction.Values[field.Name]; ok {
			continue
		}

		setEntityField(c, function, entity, alloc, field.Name, field.Value)
	}

	return &qbe.TypedValue{
		Type:  qbe.RemapAstType(entity),
		Value: alloc,
	}
}

func setEntityField(c *Compiler, function *qbe.Function, entity *ast.Entity, alloc qbe.Value, name string, expression ast.Expression) {
	structType := c.Module.GetTypeByName(entity.Name)

	offset, ty := memberToOffset(c, entity, structType, name)

	value := c.compileStatement(expression, function, ty, false)

	offsetTmp := c.createVariable(qbe.Long, "offset")
	function.LastBlock().AddAssign(offsetTmp, qbe.Long, qbe.NewAddInstruction(alloc, qbe.NewConstantValue(offset)))

	function.LastBlock().AddInstruction(qbe.NewStoreInstruction(ty, value.Value, offsetTmp))
}

func memberToOffset(c *Compiler, entity *ast.Entity, structType *qbe.TypeDefinition, field string) (int64, qbe.Type) {
	var offset int64 = 0
	var fieldType qbe.Type
	for i, item := range entity.Fields {
		if item.Name == field {
			fieldType = structType.Items[i].Type
			break
		}

		offset += int64(structType.Items[i].Type.Size(c.Module))
	}

	return offset, fieldType
}
