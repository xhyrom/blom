package resolver

import (
	"blom/ast"
	"blom/debug"
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
	"errors"
	"os"
)

type ModuleResolver struct {
	ResolvedModules map[string]*ResolvedModule
	StandardLibPath string
	Mode            ResolveMode
	SearchPaths     []string
	ProcessingStack []string
}

func NewModuleResolver(stdLibPath string, mode ResolveMode) *ModuleResolver {
	return &ModuleResolver{
		ResolvedModules: make(map[string]*ResolvedModule),
		StandardLibPath: stdLibPath,
		Mode:            mode,
		SearchPaths:     []string{".", "./lib"},
		ProcessingStack: make([]string, 0),
	}
}

func (r *ModuleResolver) AddSearchPath(path string) {
	r.SearchPaths = append(r.SearchPaths, path)
}

func (r *ModuleResolver) ResolveProgram(sourceFile string, program *ast.Program) (*ast.Program, error) {
	if r.Mode != MergeModules {
		return program, nil
	}

	result := ast.NewProgram()

	for _, node := range program.Body {
		if imp, ok := node.(*ast.Import); ok {
			err := r.processImport(sourceFile, imp, result)
			if err != nil {
				return nil, err
			}
		} else {
			result.Body = append(result.Body, node)
		}
	}

	return result, nil
}

func (r *ModuleResolver) processImport(sourceFile string, imp *ast.Import, program *ast.Program) error {
	for _, path := range r.ProcessingStack {
		if path == imp.Path {
			return errors.New("circular import detected: " + imp.Path)
		}
	}

	modulePath, err := r.findModulePath(imp.Path)
	if err != nil {
		dbg := debug.NewSourceLocationFromNode(sourceFile, imp)
		dbg.ThrowError("Cannot find module: "+imp.Path, true,
			debug.NewHint("Make sure the module exists and is in the search path", ""))
		return err
	}

	module, exists := r.ResolvedModules[imp.Path]
	if !exists {
		r.ProcessingStack = append(r.ProcessingStack, imp.Path)

		content, err := os.ReadFile(modulePath)
		if err != nil {
			return err
		}

		lex := lexer.New(modulePath, string(content))
		tkns := make([]tokens.Token, 0)

		current := lex.Next()
		for current.Kind != tokens.Eof {
			tkns = append(tkns, *current)
			current = lex.Next()
		}

		p := parser.New(modulePath)
		moduleProgram := p.AST(modulePath, string(content))

		module = &ResolvedModule{
			Path:    modulePath,
			Program: moduleProgram,
			Alias:   imp.Alias,
		}
		r.ResolvedModules[imp.Path] = module

		resolvedProgram, err := r.ResolveProgram(modulePath, moduleProgram)
		if err != nil {
			return err
		}
		module.Program = resolvedProgram

		r.ProcessingStack = r.ProcessingStack[:len(r.ProcessingStack)-1]
	}

	for _, stmt := range module.Program.Body {
		if _, isImport := stmt.(*ast.Import); !isImport {
			qualified := r.qualifyNode(stmt, imp.Path, imp.Alias)
			program.Body = append(program.Body, qualified)
		}
	}

	return nil
}
