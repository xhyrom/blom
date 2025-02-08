package types

import (
	"blom/ast"
)

func (a *TypeAnalyzer) analyzeEntityConstruction(entity *ast.EntityConstruction) ast.Type {
	entityType, _ := a.TypeManager.Get(entity.Name)
	return entityType
}
