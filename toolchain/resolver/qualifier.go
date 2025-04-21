package resolver

import (
	"blom/ast"
	"strings"
)

func (r *ModuleResolver) qualifyNode(node ast.Node, importPath string, alias string) ast.Node {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.FunctionDeclaration:
		return r.qualifyFunctionDeclaration(n, importPath, alias)
	}

	return node
}

func (r *ModuleResolver) qualifyFunctionDeclaration(fn *ast.FunctionDeclaration, importPath string, alias string) *ast.FunctionDeclaration {
	if !fn.HasAnnotation(ast.Public) {
		return fn
	}

	var prefix string
	if alias != "" {
		prefix = alias
	} else {
		parts := strings.Split(importPath, "/")
		prefix = parts[len(parts)-1]
	}

	newPath := ast.Path{
		Segments: make([]ast.IdentifierLiteral, 0),
		Loc:      fn.Loc,
	}

	newPath.Segments = append(newPath.Segments, ast.IdentifierLiteral{
		Value: prefix,
		Loc:   fn.Loc,
	})

	newPath.Segments = append(newPath.Segments, fn.Path.Segments...)

	qualified := *fn
	qualified.Path = newPath

	return &qualified
}
