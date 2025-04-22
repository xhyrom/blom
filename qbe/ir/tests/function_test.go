package tests

import (
	"testing"

	"github.com/xhyrom/blom/qbe/ir"
)

func TestFunctionString(t *testing.T) {
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

	tests := []struct {
		function ir.Function
		expected string
	}{
		{
			ir.Function{
				Linkage:    ir.NewLinkage(true),
				Name:       "foo",
				Params:     params,
				ReturnType: ir.Word,
				Variadic:   false,
				Blocks:     blocks,
			},
			"export function w $foo(w %a, w %b) {\n@start\n\t%t1 =w add %a, %b\n\tret %t1\n}",
		},
		{
			ir.Function{
				Linkage:    ir.NewLinkage(false),
				Name:       "bar",
				Params:     params,
				ReturnType: ir.Word,
				Variadic:   true,
				Blocks:     blocks,
			},
			"function w $bar(w %a, w %b, ...) {\n@start\n\t%t1 =w add %a, %b\n\tret %t1\n}",
		},
	}

	for _, test := range tests {
		if test.function.String() != test.expected {
			t.Errorf("expected %s, got %s", test.expected, test.function.String())
		}
	}
}

func TestFunctionStringNoReturnType(t *testing.T) {
	params := []ir.TypedValue{
		{Type: ir.Word, Value: ir.TemporaryValue{Name: "a"}},
		{Type: ir.Word, Value: ir.TemporaryValue{Name: "b"}},
	}

	function := ir.Function{
		Linkage:  ir.NewLinkage(true),
		Name:     "foo",
		Params:   params,
		Variadic: false,
		Blocks: []ir.Block{
			{Label: "start", Statements: []ir.Statement{
				ir.AssignStatement{Name: ir.TemporaryValue{Name: "t1"}, Type: ir.Word, Instruction: ir.AddInstruction{
					Left:  ir.TemporaryValue{Name: "a"},
					Right: ir.TemporaryValue{Name: "b"},
				}},
			}},
		},
	}

	expected := "export function $foo(w %a, w %b) {\n@start\n\t%t1 =w add %a, %b\n}"
	if function.String() != expected {
		t.Errorf("expected %s, got %s", expected, function.String())
	}
}
