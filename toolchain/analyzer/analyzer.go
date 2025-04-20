package analyzer

import (
	"blom/analyzer/manager"
	"blom/analyzer/types"
	"blom/ast"
)

type Analyzer struct {
	Source          string
	Program         *ast.Program
	FunctionManager *manager.FunctionManager
	TypeAnalyzer    *types.TypeAnalyzer
}

func New(file string, program *ast.Program) *Analyzer {
	return &Analyzer{
		Source:          file,
		Program:         program,
		FunctionManager: manager.NewFunctionManager(),
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

func (a *Analyzer) GetAnalysisContext() *types.AnalysisContext {
	return a.TypeAnalyzer.GetContext()
}

func (a *Analyzer) analyzeTypes() {
	a.TypeAnalyzer = types.New(a.Source, a.Program, a.FunctionManager)
	a.TypeAnalyzer.Analyze()
}

func (a *Analyzer) eliminateDeadCode() {
}

func (a *Analyzer) inlineFunctions() {
}

func (a *Analyzer) mergeImportedModules() {
}
