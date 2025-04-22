package compiler

import (
	"blom/analyzer"
	"blom/ast"
	"blom/compiler/codegen"
	"blom/compiler/mangling"
	"blom/compiler/resolver"
)

type Compiler struct {
	sourceFile string
	codegen    *codegen.Codegen
}

func New(sourceFile string, codegen *codegen.Codegen) *Compiler {
	return &Compiler{
		sourceFile: sourceFile,
		codegen:    codegen,
	}
}

func (c *Compiler) Compile(program *ast.Program) string {
	modules := resolver.NewModuleResolver(
		"/home/hyro/Workspace/blom/", // standard library path
		resolver.MergeModules,        // merge imported modules
	)

	program, err := modules.ResolveProgram(c.sourceFile, program)
	if err != nil {
		panic(err)
	}

	analyzer := analyzer.New(c.sourceFile, program)
	analyzer.Analyze()

	mangler := mangling.NewASTMangler()
	mangler.Mangle(program, analyzer.GetAnalysisContext())

	demangler := mangling.NewNativeDemangler()
	demangler.Demangle(program, analyzer.GetAnalysisContext())

	return c.codegen.Generate(program)
}
