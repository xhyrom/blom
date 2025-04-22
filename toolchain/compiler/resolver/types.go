package resolver

import "blom/ast"

type ResolveMode int

const (
	ParseOnly ResolveMode = iota
	MergeModules
	PreserveModules
)

type ResolvedModule struct {
	Path    string
	Program *ast.Program
	Alias   string
}
