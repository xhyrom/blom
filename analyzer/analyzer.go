package analyzer

import (
	"blom/analyzer/manager"
	"blom/analyzer/overloading"
	"blom/analyzer/types"
	"blom/ast"
)

type Analyzer struct {
	Source          string
	Program         *ast.Program
	FunctionManager *manager.FunctionManager
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

	a.overload()
	a.analyzeTypes()
	a.eliminateDeadCode()
	a.inlineFunctions()
	a.mergeImportedModules()
}

func (a *Analyzer) overload() {
	overloading.New(a.Program, a.FunctionManager).Overload()
}

func (a *Analyzer) analyzeTypes() {
	types.New(a.Source, a.Program, a.FunctionManager).Analyze()
}

func (a *Analyzer) eliminateDeadCode() {
}

func (a *Analyzer) inlineFunctions() {
}

func (a *Analyzer) mergeImportedModules() {
}
