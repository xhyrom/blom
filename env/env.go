package env

import (
	"blom/ast"
)

type Environment[T any] struct {
	Functions map[string]*ast.FunctionDeclaration
	Variables map[string]T
	Parent    *Environment[T]
}

func New[T any](environments ...Environment[T]) *Environment[T] {
	environment := Environment[T]{
		Functions: make(map[string]*ast.FunctionDeclaration),
		Variables: make(map[string]T),
		Parent:    nil,
	}

	for _, env := range environments {
		environment.Parent = &env
	}

	return &environment
}

func (env *Environment[T]) Set(key string, value T) {
	env.Variables[key] = value
}

func (env *Environment[T]) Get(key string) T {
	return env.Variables[key]
}

func (env *Environment[T]) SetFunction(key string, value *ast.FunctionDeclaration) {
	env.Functions[key] = value
}

func (env *Environment[T]) GetFunction(key string) *ast.FunctionDeclaration {
	return env.Functions[key]
}

func (env *Environment[T]) FindVariable(key string) (T, bool) {
	var zero T

	currentEnv := env
	for currentEnv != nil {
		if value, exists := currentEnv.Variables[key]; exists {
			return value, true
		}
		currentEnv = currentEnv.Parent
	}
	return zero, false
}

func (env *Environment[T]) FindFunction(key string) *ast.FunctionDeclaration {
	currentEnv := env
	for currentEnv != nil {
		if function, exists := currentEnv.Functions[key]; exists {
			return function
		}
		currentEnv = currentEnv.Parent
	}
	return nil
}

func (env *Environment[T]) Collect() {
	env.Variables = make(map[string]T)

	env.Parent = nil
}
