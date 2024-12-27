package env

import (
	"blom/ast"
	"blom/env/objects"
)

type Environment struct {
	Functions map[string]*ast.FunctionDeclaration
	Variables map[string]objects.Object
	Parent    *Environment
}

func New(environments ...Environment) *Environment {
	environment := Environment{
		Functions: make(map[string]*ast.FunctionDeclaration),
		Variables: make(map[string]objects.Object),
		Parent:    nil,
	}

	for _, env := range environments {
		environment.Parent = &env
	}

	return &environment
}

func (env *Environment) Set(key string, value objects.Object) {
	env.Variables[key] = value
}

func (env *Environment) Get(key string) objects.Object {
	return env.Variables[key]
}

func (env *Environment) SetFunction(key string, value *ast.FunctionDeclaration) {
	env.Functions[key] = value
}

func (env *Environment) GetFunction(key string) *ast.FunctionDeclaration {
	return env.Functions[key]
}

func (env *Environment) FindVariable(key string) objects.Object {
	currentEnv := env
	for currentEnv != nil {
		if value, exists := currentEnv.Variables[key]; exists {
			return value
		}
		currentEnv = currentEnv.Parent
	}
	return nil
}

func (env *Environment) FindFunction(key string) *ast.FunctionDeclaration {
	currentEnv := env
	for currentEnv != nil {
		if function, exists := currentEnv.Functions[key]; exists {
			return function
		}
		currentEnv = currentEnv.Parent
	}
	return nil
}
