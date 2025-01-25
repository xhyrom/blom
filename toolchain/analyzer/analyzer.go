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
	TypeManager     *manager.TypeManager
}

func New(file string, program *ast.Program) *Analyzer {
	return &Analyzer{
		Source:          file,
		Program:         program,
		FunctionManager: manager.NewFunctionManager(),
		TypeManager:     manager.NewTypeManager(),
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
	types.New(a.Source, a.Program, a.FunctionManager, a.TypeManager).Analyze()
}

func (a *Analyzer) eliminateDeadCode() {
}

func (a *Analyzer) inlineFunctions() {
}

func (a *Analyzer) mergeImportedModules() {
}
