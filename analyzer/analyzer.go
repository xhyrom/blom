package analyzer

import (
	"blom/analyzer/types"
	"blom/ast"
)

type Analyzer struct {
	Source  string
	Program *ast.Program
}

func New(file string, program *ast.Program) *Analyzer {
	return &Analyzer{
		Source:  file,
		Program: program,
	}
}

func (a *Analyzer) Analyze() {
	a.analyzeTypes()
	a.eliminateDeadCode()
	a.inlineFunctions()
	a.mergeImportedModules()
}

func (a *Analyzer) analyzeTypes() {
	types.New(a.Source, a.Program).Analyze()
}

func (a *Analyzer) eliminateDeadCode() {
}

func (a *Analyzer) inlineFunctions() {
}

func (a *Analyzer) mergeImportedModules() {
}
