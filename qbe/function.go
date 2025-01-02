package qbe

import (
	"blom/ast"
	"fmt"
	"strings"
)

type Linkage struct {
	Exported bool
	Section  *string
	Secflags *string
}

func NewLinkage(exported bool) Linkage {
	return Linkage{
		Exported: exported,
	}
}

func (l Linkage) String() string {
	result := ""

	if l.Exported {
		result += "export "
	}

	if l.Section != nil {
		result += fmt.Sprintf("section \"%s\"", *l.Section)

		if l.Secflags != nil {
			result += fmt.Sprintf(" \"%s\"", *l.Secflags)
		}

		result += " "
	}

	return result
}

type Function struct {
	Linkage    Linkage
	Name       string
	Arguments  []TypedValue
	ReturnType Type
	Variadic   bool
	External   bool
	Blocks     []Block
}

func (f *Function) AddBlock(label string) {
	f.Blocks = append(f.Blocks, Block{
		Label:      label,
		Statements: make([]Statement, 0),
	})
}

func (f *Function) LastBlock() *Block {
	return &f.Blocks[len(f.Blocks)-1]
}

func (f Function) String() string {
	signature := fmt.Sprintf("%sfunction", f.Linkage)

	if f.ReturnType != nil {
		signature += fmt.Sprintf(" %s", f.ReturnType.IntoAbi())
	}

	parameters := make([]string, len(f.Arguments))

	for i, p := range f.Arguments {
		parameters[i] = fmt.Sprintf("%s %s", p.Type, p.Value)
	}

	if f.Variadic {
		parameters = append(parameters, "...")
	}

	signature += fmt.Sprintf(" $%s(%s)", f.Name, strings.Join(parameters, ", "))

	blocks := make([]string, len(f.Blocks))

	for i, b := range f.Blocks {
		blocks[i] = fmt.Sprintf("%s", b)
	}

	return fmt.Sprintf("%s {\n%s\n}", signature, strings.Join(blocks, "\n"))
}

func RemapAstFunction(fun ast.FunctionDeclaration) Function {
	arguments := make([]TypedValue, len(fun.Arguments))

	for i, arg := range fun.Arguments {
		arguments[i] = TypedValue{
			Type:  RemapAstType(arg.Type),
			Value: NewTemporaryValue(arg.Name),
		}
	}

	return Function{
		Linkage:    NewLinkage(fun.HasAnnotation(ast.Public)),
		Name:       fun.Name,
		Arguments:  arguments,
		ReturnType: RemapAstType(fun.ReturnType),
		Variadic:   fun.Variadic,
		External:   fun.IsNative(),
		Blocks:     make([]Block, 0),
	}
}
