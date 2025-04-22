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
	case *ast.Return:
		n.Value = r.qualifyExpression(n.Value, importPath, alias)
		return n
	case *ast.Block:
		for i, stmt := range n.Body {
			n.Body[i] = r.qualifyNode(stmt, importPath, alias)
		}
		return n
	case *ast.If:
		n.Condition = r.qualifyExpression(n.Condition, importPath, alias)
		n.Then = r.qualifyNode(n.Then, importPath, alias).(*ast.Block)
		if n.Else != nil {
			n.Else = r.qualifyNode(n.Else, importPath, alias).(*ast.Block)
		}
		return n
	case *ast.Assignment:
		n.Right = r.qualifyExpression(n.Right, importPath, alias)
		return n
	case *ast.VariableDeclaration:
		if n.Init != nil {
			n.Init = r.qualifyExpression(n.Init, importPath, alias)
		}
		return n
	case *ast.WhileLoop:
		n.Condition = r.qualifyExpression(n.Condition, importPath, alias)
		n.Block = r.qualifyNode(n.Block, importPath, alias).(*ast.Block)
		return n
	}

	return node
}

func (r *ModuleResolver) qualifyExpression(expr ast.Node, importPath string, alias string) ast.Node {
	if expr == nil {
		return nil
	}

	switch e := expr.(type) {
	case *ast.Path:
		return r.qualifyPath(e, importPath, alias)
	case *ast.FunctionCall:
		e.Path = *r.qualifyPath(&e.Path, importPath, alias)
		for i, arg := range e.Args {
			e.Args[i] = r.qualifyExpression(arg, importPath, alias)
		}
		return e
	case *ast.MethodCall:
		e.Callee = r.qualifyExpression(e.Callee, importPath, alias).(*ast.Field)
		for i, arg := range e.Args {
			e.Args[i] = r.qualifyExpression(arg, importPath, alias)
		}
		return e
	case *ast.BinaryExpression:
		e.Left = r.qualifyExpression(e.Left, importPath, alias)
		e.Right = r.qualifyExpression(e.Right, importPath, alias)
		return e
	case *ast.UnaryExpression:
		e.Operand = r.qualifyExpression(e.Operand, importPath, alias)
		return e
	case *ast.Field:
		e.Base = r.qualifyExpression(e.Base, importPath, alias)
		e.Member = r.qualifyExpression(e.Member, importPath, alias)
		return e
	}

	return expr
}

func (r *ModuleResolver) qualifyPath(path *ast.Path, importPath string, alias string) *ast.Path {
	if len(path.Segments) > 1 {
		var prefix string
		if alias != "" {
			prefix = alias
		} else {
			parts := strings.Split(importPath, "/")
			prefix = parts[len(parts)-1]
		}

		newPath := ast.Path{
			Segments: make([]ast.IdentifierLiteral, 0),
			Loc:      path.Loc,
		}

		newPath.Segments = append(newPath.Segments, ast.IdentifierLiteral{
			Value: prefix,
			Loc:   path.Loc,
		})

		newPath.Segments = append(newPath.Segments, path.Segments...)

		return &newPath
	}

	return path
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

	if qualified.Block != nil {
		qualified.Block = r.qualifyNode(qualified.Block, importPath, alias).(*ast.Block)
	}

	return &qualified
}
