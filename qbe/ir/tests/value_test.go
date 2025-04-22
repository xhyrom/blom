package tests

import (
	"testing"

	"github.com/xhyrom/blom/qbe/ir"
)

func TestTemporaryValue(t *testing.T) {
	v := ir.TemporaryValue{Name: "name"}

	if v.Type() != ir.TemporaryValueType {
		t.Errorf("Expected TemporaryValueType, got %v", v.Type())
	}

	if v.String() != "%name" {
		t.Errorf("Expected %%name, got %v", v.String())
	}
}

func TestGlobalValue(t *testing.T) {
	v := ir.GlobalValue{Name: "name"}

	if v.Type() != ir.GlobalValueType {
		t.Errorf("Expected GlobalValueType, got %v", v.Type())
	}

	if v.String() != "$name" {
		t.Errorf("Expected $name, got %v", v.String())
	}
}

func TestTypedValue(t *testing.T) {
	v := ir.TypedValue{
		Value: ir.TemporaryValue{Name: "name"},
		Type:  ir.Word,
	}

	if v.String() != "w %name" {
		t.Errorf("Expected w %%name, got %v", v.String())
	}

	if v.AbiString() != "w %name" {
		t.Errorf("Expected w %%name, got %v", v.AbiString())
	}
}

func TestTypedValueDifferentAbi(t *testing.T) {
	v := ir.TypedValue{
		Value: ir.TemporaryValue{Name: "name"},
		Type:  ir.Byte,
	}

	if v.String() != "b %name" {
		t.Errorf("Expected b %%name, got %v", v.String())
	}

	if v.AbiString() != "w %name" {
		t.Errorf("Expected w %%name, got %v", v.AbiString())
	}
}
