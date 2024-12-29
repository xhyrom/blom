package analyzer

import (
	"blom/analyzer/types"
	"blom/ast"
)

type Analyzer struct {
	Source    string
	Program   *ast.Program
	Functions map[string]*ast.FunctionDeclaration
}

func New(file string, program *ast.Program) *Analyzer {
	return &Analyzer{
		Source:    file,
		Program:   program,
		Functions: make(map[string]*ast.FunctionDeclaration),
	}
}

func (a *Analyzer) Analyze() {
	// populator
	a.populate()

	a.analyzeTypes()
	a.eliminateDeadCode()
	a.inlineFunctions()
	a.mergeImportedModules()
}

func (a *Analyzer) analyzeTypes() {
	types.New(a.Source, a.Program, a.Functions).Analyze()
}

func (a *Analyzer) eliminateDeadCode() {
}

func (a *Analyzer) inlineFunctions() {
}

func (a *Analyzer) mergeImportedModules() {
}
