package tests

import (
	"blom/qbe"
	"testing"
)

func TestFunctionString(t *testing.T) {
	returnType := qbe.Word
	arguments := []qbe.TypedValue{
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

	tests := []struct {
		function qbe.Function
		expected string
	}{
		{
			qbe.Function{
				Linkage:    qbe.NewLinkage(true),
				Name:       "foo",
				Arguments:  arguments,
				ReturnType: &returnType,
				Variadic:   false,
				Blocks:     blocks,
			},
			"export function w $foo(w %a, w %b) {\n@start\n\t%t1 =w add %a, %b\n\tret %t1\n}",
		},
		{
			qbe.Function{
				Linkage:    qbe.NewLinkage(false),
				Name:       "bar",
				Arguments:  arguments,
				ReturnType: &returnType,
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
	arguments := []qbe.TypedValue{
		{Type: qbe.Word, Value: qbe.TemporaryValue{Name: "a"}},
		{Type: qbe.Word, Value: qbe.TemporaryValue{Name: "b"}},
	}

	function := qbe.Function{
		Linkage:   qbe.NewLinkage(true),
		Name:      "foo",
		Arguments: arguments,
		Variadic:  false,
		Blocks: []qbe.Block{
			{Label: "start", Statements: []qbe.Statement{
				qbe.AssignStatement{Name: qbe.TemporaryValue{Name: "t1"}, Type: qbe.Word, Instruction: qbe.AddInstruction{
					Left:  qbe.TemporaryValue{Name: "a"},
					Right: qbe.TemporaryValue{Name: "b"},
				}},
			}},
		},
	}

	expected := "export function $foo(w %a, w %b) {\n@start\n\t%t1 =w add %a, %b\n}"
	if function.String() != expected {
		t.Errorf("expected %s, got %s", expected, function.String())
	}
}
