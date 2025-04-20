package mangling

import (
	"blom/analyzer/types"
	"blom/ast"
)

type FunctionPtrMap map[*ast.FunctionDeclaration]string
type FunctionNameMap map[string]string

func Mangle(program *ast.Program, analysisContext *types.AnalysisContext) {
	ptrMap, nameMap := buildFunctionMaps(program)

	for _, node := range program.Body {
		mangleTopLevelNode(node, ptrMap, nameMap, analysisContext)
	}
}

func buildFunctionMaps(program *ast.Program) (FunctionPtrMap, FunctionNameMap) {
	ptrMap := make(FunctionPtrMap)
	nameMap := make(FunctionNameMap)

	for _, node := range program.Body {
		if fn, ok := node.(*ast.FunctionDeclaration); ok {
			if shouldSkipMangling(fn) {
				continue
			}

			mangledName := GenerateFunctionMangledName(fn)

			ptrMap[fn] = mangledName

			nameMap[getFunctionKey(fn)] = mangledName
		}
	}

	return ptrMap, nameMap
}

func mangleTopLevelNode(node ast.Node, ptrMap FunctionPtrMap, nameMap FunctionNameMap, analysisContext *types.AnalysisContext) {
	if node == nil {
		return
	}

	if fn, ok := node.(*ast.FunctionDeclaration); ok {
		mangleFunctionDeclaration(fn, ptrMap)

		if fn.Block != nil {
			for _, stmt := range fn.Block.Body {
				mangleNode(stmt, ptrMap, nameMap, analysisContext)
			}
		}
	}
}

func mangleNode(node ast.Node, ptrMap FunctionPtrMap, nameMap FunctionNameMap, analysisContext *types.AnalysisContext) {
	if node == nil {
		return
	}

	switch n := node.(type) {
	case *ast.FunctionCall, *ast.MethodCall, *ast.InfixCall:
		mangleCall(n.(ast.Call), ptrMap, nameMap, analysisContext)
	case *ast.Block:
		for _, stmt := range n.Body {
			mangleNode(stmt, ptrMap, nameMap, analysisContext)
		}
	case *ast.Return:
		if n.Value != nil {
			mangleNode(n.Value, ptrMap, nameMap, analysisContext)
		}
	case *ast.If:
		mangleNode(n.Condition, ptrMap, nameMap, analysisContext)

		for _, stmt := range n.Then.Body {
			mangleNode(stmt, ptrMap, nameMap, analysisContext)
		}

		if n.Else != nil {
			for _, stmt := range n.Else.Body {
				mangleNode(stmt, ptrMap, nameMap, analysisContext)
			}
		}
	case *ast.BinaryExpression:
		mangleNode(n.Left, ptrMap, nameMap, analysisContext)
		mangleNode(n.Right, ptrMap, nameMap, analysisContext)
	case *ast.UnaryExpression:
		mangleNode(n.Operand, ptrMap, nameMap, analysisContext)
	case *ast.WhileLoop:
		mangleNode(n.Condition, ptrMap, nameMap, analysisContext)

		for _, stmt := range n.Block.Body {
			mangleNode(stmt, ptrMap, nameMap, analysisContext)
		}
	case *ast.VariableDeclaration:
		if n.Init != nil {
			mangleNode(n.Init, ptrMap, nameMap, analysisContext)
		}
	case *ast.Assignment:
		mangleNode(n.Left, ptrMap, nameMap, analysisContext)
		mangleNode(n.Right, ptrMap, nameMap, analysisContext)
	}
}
