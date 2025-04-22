package tests

import (
	"testing"

	"github.com/xhyrom/blom/qbe/ir"
)

func TestModuleString(t *testing.T) {
	align := uint64(8)
	params := []ir.TypedValue{
		{Type: ir.Word, Value: ir.TemporaryValue{Name: "a"}},
		{Type: ir.Word, Value: ir.TemporaryValue{Name: "b"}},
	}
	blocks := []ir.Block{
		{Label: "start", Statements: []ir.Statement{
			ir.AssignStatement{Name: ir.TemporaryValue{Name: "t1"}, Type: ir.Word, Instruction: ir.AddInstruction{
				Left:  ir.TemporaryValue{Name: "a"},
				Right: ir.TemporaryValue{Name: "b"},
			}},
			ir.VolatileStatement{Instruction: ir.ReturnInstruction{
				Value: ir.TemporaryValue{Name: "t1"},
			}},
		}},
	}
	functions := []ir.Function{
		{
			Linkage:    ir.NewLinkage(true),
			Name:       "foo",
			Params:     params,
			ReturnType: ir.Word,
			Variadic:   false,
			Blocks:     blocks,
		},
	}
	types := []ir.TypeDefinition{
		{Name: "myType", Items: []ir.TypedTypeDefinitionItem{
			{
				Count: 1,
				Type:  ir.Word,
			},
			{
				Count: 1,
				Type:  ir.Word,
			},
		}},
	}
	data := []ir.Data{
		{
			Linkage: ir.NewLinkage(true),
			Name:    "myData",
			Align:   &align,
			Items: []ir.TypedDataItem{
				{Item: ir.StringDataItem{Value: "hello"}, Type: ir.NewPointer(ir.Char)},
				{Item: ir.ConstantDataItem{Value: 42}, Type: ir.Long},
			},
		},
	}

	tests := []struct {
		module   ir.Module
		expected string
	}{
		{
			ir.Module{
				Functions: functions,
				Types:     types,
				Data:      data,
			},
			"type :myType = { w, w }\nexport data $myData = align 8 { l \"hello\", l 42 }\nexport function w $foo(w %a, w %b) {\n@start\n\t%t1 =w add %a, %b\n\tret %t1\n}",
		},
		{
			ir.Module{
				Functions: []ir.Function{},
				Types:     []ir.TypeDefinition{},
				Data:      []ir.Data{},
			},
			"",
		},
	}

	for _, test := range tests {
		if test.module.String() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, test.module.String())
		}
	}
}
