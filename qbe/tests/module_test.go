package tests

import (
	"blom/qbe"
	"testing"
)

func TestModuleString(t *testing.T) {
	align := uint64(8)
	returnType := qbe.Word
	parameters := []qbe.TypedValue{
		{Type: qbe.Word, Value: qbe.TemporaryValue{Name: "a"}},
		{Type: qbe.Word, Value: qbe.TemporaryValue{Name: "b"}},
	}
	blocks := []qbe.Block{
		{Label: "start", Statements: []qbe.Statement{
			qbe.AssignStatement{Name: qbe.TemporaryValue{Name: "t1"}, Type: qbe.Word, Instruction: qbe.AddInstruction{
				Left:  qbe.TemporaryValue{Name: "a"},
				Right: qbe.TemporaryValue{Name: "b"},
			}},
			qbe.VolatileStatement{Instruction: qbe.ReturnInstruction{
				Value: qbe.TemporaryValue{Name: "t1"},
			}},
		}},
	}
	functions := []qbe.Function{
		{
			Linkage:    qbe.NewLinkage(true),
			Name:       "foo",
			Parameters: parameters,
			ReturnType: &returnType,
			Variadic:   false,
			Blocks:     blocks,
		},
	}
	types := []qbe.TypeDefinition{
		{Name: "myType", Items: []qbe.TypedTypeDefinitionItem{
			{
				Count: 1,
				Type:  qbe.Word,
			},
			{
				Count: 1,
				Type:  qbe.Word,
			},
		}},
	}
	data := []qbe.Data{
		{
			Linkage: qbe.NewLinkage(true),
			Name:    "myData",
			Align:   &align,
			Items: []qbe.TypedDataItem{
				{Item: qbe.StringDataItem{Value: "hello"}, Type: qbe.String},
				{Item: qbe.ConstantDataItem{Value: 42}, Type: qbe.Long},
			},
		},
	}

	tests := []struct {
		module   qbe.Module
		expected string
	}{
		{
			qbe.Module{
				Functions: functions,
				Types:     types,
				Data:      data,
			},
			"type :myType = { w, w }\nexported data myData = align 8 { l \"hello\", l 42 }\nexported function w foo(w %a, w %b) {\n@start\n\t%t1 =w add %a, %b\n\tret %t1\n}",
		},
		{
			qbe.Module{
				Functions: []qbe.Function{},
				Types:     []qbe.TypeDefinition{},
				Data:      []qbe.Data{},
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
