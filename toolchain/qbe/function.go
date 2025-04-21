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
	Params     []TypedValue
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

func (f *Function) FirstBlock() *Block {
	return &f.Blocks[0]
}

func (f *Function) LastBlock() *Block {
	return &f.Blocks[len(f.Blocks)-1]
}

func (f Function) String() string {
	signature := fmt.Sprintf("%sfunction", f.Linkage)

	if f.ReturnType != nil {
		signature += fmt.Sprintf(" %s", f.ReturnType.IntoAbi())
	}

	parameters := make([]string, len(f.Params))

	for i, p := range f.Params {
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
	params := make([]TypedValue, len(fun.Params))

	for i, param := range fun.Params {
		params[i] = TypedValue{
			Type:  RemapAstType(param.Type),
			Value: NewTemporaryValue(param.Name.Value),
		}
	}

	return Function{
		Linkage:    NewLinkage(fun.HasAnnotation(ast.Public)),
		Name:       fun.Path.Dotify(),
		Params:     params,
		ReturnType: RemapAstType(fun.Return),
		Variadic:   fun.HasAnnotation(ast.Variadic),
		External:   fun.HasAnnotation(ast.Native),
		Blocks:     make([]Block, 0),
	}
}

// FunctionBox is a wrapped around a Type that represents a function
// It holds the function
type FunctionBox struct {
	Inner Function
}

func NewFunctionBox(inner Function) FunctionBox {
	return FunctionBox{Inner: inner}
}

func (f FunctionBox) String() string {
	return Func.String()
}

func (f FunctionBox) IsNumeric() bool {
	return Func.IsNumeric()
}

func (f FunctionBox) IsInteger() bool {
	return Func.IsInteger()
}

func (f FunctionBox) IsFloatingPoint() bool {
	return Func.IsFloatingPoint()
}

func (f FunctionBox) IsSigned() bool {
	return Func.IsSigned()
}

func (f FunctionBox) IsUnsigned() bool {
	return Func.IsUnsigned()
}

func (f FunctionBox) IsPointer() bool {
	return Func.IsPointer()
}

func (f FunctionBox) IsFunction() bool {
	return Func.IsFunction()
}

func (f FunctionBox) IsStruct() bool {
	return Func.IsStruct()
}

func (f FunctionBox) IsMapToInt() bool {
	return Func.IsMapToInt()
}

func (f FunctionBox) Weight() uint8 {
	return Func.Weight()
}

func (f FunctionBox) Size(module Module) uint64 {
	return Func.Size(module)
}

func (f FunctionBox) IntoAbi() Type {
	return Func.IntoAbi()
}

func (f FunctionBox) IntoBase() Type {
	return Func.IntoBase()
}
