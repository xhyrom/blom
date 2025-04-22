package qbe

import (
	"blom/ast"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileFunction(declaration *ast.FunctionDeclaration) {
	c.Scopes.Append()

	var linkage ir.Linkage
	if declaration.HasAnnotation(ast.Public) {
		linkage = ir.NewLinkage(true)
	} else {
		linkage = ir.NewLinkage(false)
	}

	returnType := RemapAstType(declaration.Return)
	function := ir.Function{
		Linkage:    linkage,
		Name:       declaration.Path.Dotify(),
		Params:     make([]ir.TypedValue, len(declaration.Params)),
		ReturnType: returnType,
		Variadic:   declaration.Variadic,
		External:   declaration.HasAnnotation(ast.Native),
		Blocks:     make([]ir.Block, 0),
	}

	if declaration.HasAnnotation(ast.Native) {
		c.Scopes.Pop()
		return
	}

	function.AddBlock("start")

	for i, param := range declaration.Params {
		t := RemapAstType(param.Type)

		temp := c.createVariable(t, param.Name.Value)
		stmt := &ast.VariableDeclaration{
			Name: param.Name,
			Type: param.Type,
			Init: &ast.IdentifierLiteral{
				Value: param.Name.Value,
			},
		}

		c.compileStatement(stmt, &function, t, false)
		function.Params[i] = ir.NewTypedValue(t, temp)
	}

	for _, statement := range declaration.Block.Body {
		c.compileStatement(statement, &function, nil, false)
	}

	c.Scopes.Pop()

	c.Module.SetFunctionByName(declaration.Path.Dotify(), function)
}
