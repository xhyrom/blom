package env

import "blom/ast"

type Environment struct {
	Functions map[string]*ast.FunctionDeclaration
	Variables map[string]Object
}

func New(environments ...Environment) *Environment {
	environment := Environment{
		Functions: make(map[string]*ast.FunctionDeclaration),
		Variables: make(map[string]Object),
	}

	for _, env := range environments {
		for key, value := range env.Variables {
			environment.Variables[key] = value
		}

		for key, value := range env.Functions {
			environment.Functions[key] = value
		}
	}

	return &environment
}

func (env *Environment) Set(key string, value Object) {
	env.Variables[key] = value
}

func (env *Environment) Get(key string) Object {
	return env.Variables[key]
}

func (env *Environment) SetFunction(key string, value *ast.FunctionDeclaration) {
	env.Functions[key] = value
}

func (env *Environment) GetFunction(key string) *ast.FunctionDeclaration {
	return env.Functions[key]
}
