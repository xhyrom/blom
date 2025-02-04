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
	alloc := c.createVariable(qbe.Long, construction.Name)

	function.LastBlock().AddAssign(alloc, qbe.Long, qbe.NewAlloc8Instruction(qbe.NewConstantValue(4)))

	offset := c.createVariable(qbe.Long, "offset")
	function.LastBlock().AddAssign(offset, qbe.Long, qbe.NewAddInstruction(alloc, qbe.NewConstantValue(0)))

	function.LastBlock().AddInstruction(qbe.NewStoreInstruction(qbe.Word, qbe.NewConstantValue(100), offset))

	return &qbe.TypedValue{
		Type:  vtype,
		Value: alloc,
	}
}
