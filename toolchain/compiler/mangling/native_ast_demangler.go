package mangling

import (
	"blom/analyzer/types"
	"blom/ast"
)

type NativeDemangler struct {
	functionCallMap types.FunctionCallMap
}

func NewNativeDemangler() *NativeDemangler {
	return &NativeDemangler{}
}

func (d *NativeDemangler) Demangle(program *ast.Program, analysisContext *types.AnalysisContext) {
	d.functionCallMap = analysisContext.FunctionCalls

	for _, node := range program.Body {
		d.demangleTopLevelNode(node)
	}

	for _, node := range program.Body {
		if fn, ok := node.(*ast.FunctionDeclaration); ok {
			d.demangleFunctionDeclaration(fn)
		}
	}
}

func (d *NativeDemangler) demangleFunctionDeclaration(fn *ast.FunctionDeclaration) {
	if fn.HasAnnotation(ast.Native) {
		if len(fn.Path.Segments) > 1 {
			lastSegment := fn.Path.Segments[len(fn.Path.Segments)-1]

			fn.Path = ast.Path{
				Segments: []ast.IdentifierLiteral{lastSegment},
				Loc:      fn.Path.Loc,
			}
		}
	}
}

func (d *NativeDemangler) demangleTopLevelNode(node ast.Node) {
	if node == nil {
		return
	}

	if fn, ok := node.(*ast.FunctionDeclaration); ok {
		if fn.Block != nil {
			for _, stmt := range fn.Block.Body {
				d.demangleNode(stmt)
			}
		}
	}
}

func (d *NativeDemangler) demangleNode(node ast.Node) {
	if node == nil {
		return
	}

	switch n := node.(type) {
	case *ast.FunctionCall:
		if n == nil {
			return
		}

		d.demangleFunctionCall(n)
	case *ast.MethodCall:
		if n == nil {
			return
		}

		d.demangleMethodCall(n)
	case *ast.InfixCall:
		if n == nil {
			return
		}

		d.demangleInfixCall(n)
	case *ast.Block:
		if n == nil {
			return
		}

		for _, stmt := range n.Body {
			d.demangleNode(stmt)
		}
	case *ast.Return:
		if n == nil {
			return
		}

		d.demangleNode(n.Value)
	case *ast.If:
		if n == nil {
			return
		}

		d.demangleNode(n.Condition)
		d.demangleNode(n.Then)
		d.demangleNode(n.Else)
	case *ast.BinaryExpression:
		if n == nil {
			return
		}

		d.demangleNode(n.Left)
		d.demangleNode(n.Right)
	case *ast.UnaryExpression:
		if n == nil {
			return
		}

		d.demangleNode(n.Operand)
	case *ast.WhileLoop:
		if n == nil {
			return
		}

		d.demangleNode(n.Condition)
		d.demangleNode(n.Block)
	case *ast.VariableDeclaration:
		if n == nil {
			return
		}

		d.demangleNode(n.Init)
	case *ast.Assignment:
		if n == nil {
			return
		}

		d.demangleNode(n.Left)
		d.demangleNode(n.Right)
	}
}

func (d *NativeDemangler) demangleFunctionCall(call *ast.FunctionCall) {
	for _, arg := range call.Args {
		d.demangleNode(arg)
	}

	targetFn, exists := d.functionCallMap[call]
	if !exists {
		return
	}

	if targetFn.HasAnnotation(ast.Native) {
		if len(call.Path.Segments) > 1 {
			lastSegment := call.Path.Segments[len(call.Path.Segments)-1]

			call.Path = ast.Path{
				Segments: []ast.IdentifierLiteral{lastSegment},
				Loc:      call.Path.Loc,
			}
		}
	}
}

func (d *NativeDemangler) demangleMethodCall(call *ast.MethodCall) {
	d.demangleNode(call.Callee)
	for _, arg := range call.Args {
		d.demangleNode(arg)
	}
}

func (d *NativeDemangler) demangleInfixCall(call *ast.InfixCall) {
	if call.FunctionCall != nil {
		d.demangleFunctionCall(call.FunctionCall)
	} else if call.MethodCall != nil {
		d.demangleMethodCall(call.MethodCall)
	}
}
