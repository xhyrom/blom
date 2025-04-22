package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileFunction(declaration *ast.FunctionDeclaration) {
	c.Scopes.Append()

	var linkage qbe.Linkage
	if declaration.HasAnnotation(ast.Public) {
		linkage = qbe.NewLinkage(true)
	} else {
		linkage = qbe.NewLinkage(false)
	}

	returnType := qbe.RemapAstType(declaration.Return)
	function := qbe.Function{
		Linkage:    linkage,
		Name:       declaration.Path.Dotify(),
		Params:     make([]qbe.TypedValue, len(declaration.Params)),
		ReturnType: returnType,
		Variadic:   declaration.Variadic,
		External:   declaration.HasAnnotation(ast.Native),
		Blocks:     make([]qbe.Block, 0),
	}

	if declaration.HasAnnotation(ast.Native) {
		c.Scopes.Pop()
		return
	}

	function.AddBlock("start")

	for i, param := range declaration.Params {
		t := qbe.RemapAstType(param.Type)

		temp := c.createVariable(t, param.Name.Value)
		stmt := &ast.VariableDeclaration{
			Name: param.Name,
			Type: param.Type,
			Init: &ast.IdentifierLiteral{
				Value: param.Name.Value,
			},
		}

		c.compileStatement(stmt, &function, t, false)
		function.Params[i] = qbe.NewTypedValue(t, temp)
	}

	for _, statement := range declaration.Block.Body {
		c.compileStatement(statement, &function, nil, false)
	}

	c.Scopes.Pop()

	c.Module.SetFunctionByName(declaration.Path.Dotify(), function)
}
