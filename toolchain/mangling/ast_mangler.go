package mangling

import (
	"blom/analyzer/types"
	"blom/ast"
)

type ASTMangler struct {
	encoder         *Encoder
	functionNameMap map[*ast.FunctionDeclaration]string
}

func NewASTMangler() *ASTMangler {
	return &ASTMangler{
		encoder:         NewEncoder(),
		functionNameMap: make(map[*ast.FunctionDeclaration]string),
	}
}

func (m *ASTMangler) Mangle(program *ast.Program, analysisContext *types.AnalysisContext) {
	m.buildFunctionMaps(program)

	for _, node := range program.Body {
		m.mangleTopLevelNode(node, analysisContext)
	}
}

func (m *ASTMangler) buildFunctionMaps(program *ast.Program) {
	for _, node := range program.Body {
		if fn, ok := node.(*ast.FunctionDeclaration); ok {
			if shouldSkipMangling(fn) {
				continue
			}

			paramTypes := make([]ast.Type, len(fn.Params))
			for i, param := range fn.Params {
				paramTypes[i] = param.Type
			}

			mangledName := m.encoder.EncodeFunctionName(fn.Path.Dotify(), paramTypes, fn.Return)
			m.functionNameMap[fn] = string(mangledName)
		}
	}
}

func (m *ASTMangler) mangleTopLevelNode(node ast.Node, analysisContext *types.AnalysisContext) {
	if node == nil {
		return
	}

	if fn, ok := node.(*ast.FunctionDeclaration); ok {
		m.mangleFunctionDeclaration(fn)

		if fn.Block != nil {
			for _, stmt := range fn.Block.Body {
				m.mangleNode(stmt, analysisContext)
			}
		}
	}
}

func (m *ASTMangler) mangleNode(node ast.Node, analysisContext *types.AnalysisContext) {
	if node == nil {
		return
	}

	switch n := node.(type) {
	case *ast.FunctionCall, *ast.MethodCall, *ast.InfixCall:
		m.mangleCall(n.(ast.Call), analysisContext)
	case *ast.Block:
		for _, stmt := range n.Body {
			m.mangleNode(stmt, analysisContext)
		}
	case *ast.Return:
		m.mangleNode(n.Value, analysisContext)
	case *ast.If:
		m.mangleNode(n.Condition, analysisContext)
		m.mangleNode(n.Then, analysisContext)
		m.mangleNode(n.Else, analysisContext)
	case *ast.BinaryExpression:
		m.mangleNode(n.Left, analysisContext)
		m.mangleNode(n.Right, analysisContext)
	case *ast.UnaryExpression:
		m.mangleNode(n.Operand, analysisContext)
	case *ast.WhileLoop:
		m.mangleNode(n.Condition, analysisContext)
		m.mangleNode(n.Block, analysisContext)
	case *ast.VariableDeclaration:
		m.mangleNode(n.Init, analysisContext)
	case *ast.Assignment:
		m.mangleNode(n.Left, analysisContext)
		m.mangleNode(n.Right, analysisContext)
	}
}

func (m *ASTMangler) mangleFunctionDeclaration(fn *ast.FunctionDeclaration) {
	if shouldSkipMangling(fn) {
		return
	}

	if mangledName, exists := m.functionNameMap[fn]; exists {
		segments := []ast.IdentifierLiteral{
			{
				Value: mangledName,
				Loc:   fn.Loc,
			},
		}

		fn.Path.Segments = segments
	}
}

func (m *ASTMangler) mangleCall(call ast.Call, analysisContext *types.AnalysisContext) {
	switch c := call.(type) {
	case *ast.FunctionCall:
		m.mangleFunctionCall(c, analysisContext)
	case *ast.MethodCall:
		m.mangleMethodCall(c, analysisContext)
	case *ast.InfixCall:
		m.mangleInfixCall(c, analysisContext)
	default:
		panic("unknown call type")
	}
}

func (m *ASTMangler) mangleFunctionCall(call *ast.FunctionCall, analysisContext *types.AnalysisContext) {
	for _, arg := range call.Args {
		m.mangleNode(arg, analysisContext)
	}

	if fn, exists := analysisContext.FunctionCalls[call]; exists {
		if !shouldSkipMangling(fn) {
			if mangledName, exists := m.functionNameMap[fn]; exists {
				segments := []ast.IdentifierLiteral{
					{
						Value: mangledName,
						Loc:   call.Path.Loc,
					},
				}
				call.Path.Segments = segments
			}
		}
	}
}

func (m *ASTMangler) mangleMethodCall(call *ast.MethodCall, analysisContext *types.AnalysisContext) {
	panic("not implemented method call mangling")
}

func (m *ASTMangler) mangleInfixCall(call *ast.InfixCall, analysisContext *types.AnalysisContext) {
	if call.FunctionCall != nil {
		m.mangleFunctionCall(call.FunctionCall, analysisContext)
		return
	}

	m.mangleMethodCall(call.MethodCall, analysisContext)
}

func updateCallPath(call *ast.FunctionCall, path ast.Path) {
	segments := make([]ast.IdentifierLiteral, len(path.Segments))

	for i, part := range path.Segments {
		segments[i] = part
	}

	call.Path.Segments = segments
}
