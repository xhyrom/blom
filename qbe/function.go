package qbe

import (
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
		result += "exported "
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
	Parameters []TypedValue
	ReturnType *Type
	Variadic   bool
	Blocks     []Block
}

func (f Function) String() string {
	signature := fmt.Sprintf("%sfunction", f.Linkage)

	if f.ReturnType != nil {
		signature += fmt.Sprintf(" %s", (*f.ReturnType).IntoAbi())
	}

	parameters := make([]string, len(f.Parameters))

	for i, p := range f.Parameters {
		parameters[i] = fmt.Sprintf("%s %s", p.Type, p.Value)
	}

	if f.Variadic {
		parameters = append(parameters, "...")
	}

	signature += fmt.Sprintf(" %s(%s)", f.Name, strings.Join(parameters, ", "))

	blocks := make([]string, len(f.Blocks))

	for i, b := range f.Blocks {
		blocks[i] = fmt.Sprintf("%s", b)
	}

	return fmt.Sprintf("%s {\n%s\n}", signature, strings.Join(blocks, "\n"))
}
