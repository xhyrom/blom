package tests

import (
	"blom/qbe"
	"testing"
)

func TestTemporaryValue(t *testing.T) {
	v := qbe.TemporaryValue{Name: "name"}

	if v.Type() != qbe.TemporaryValueType {
		t.Errorf("Expected TemporaryValueType, got %v", v.Type())
	}

	if v.String() != "%name" {
		t.Errorf("Expected %%name, got %v", v.String())
	}
}

func TestGlobalValue(t *testing.T) {
	v := qbe.GlobalValue{Name: "name"}

	if v.Type() != qbe.GlobalValueType {
		t.Errorf("Expected GlobalValueType, got %v", v.Type())
	}

	if v.String() != "$name" {
		t.Errorf("Expected $name, got %v", v.String())
	}
}

func TestTypedValue(t *testing.T) {
	v := qbe.TypedValue{
		Value: qbe.TemporaryValue{Name: "name"},
		Type:  qbe.Word,
	}

	if v.String() != "w %name" {
		t.Errorf("Expected w %%name, got %v", v.String())
	}

	if v.AbiString() != "w %name" {
		t.Errorf("Expected w %%name, got %v", v.AbiString())
	}
}

func TestTypedValueDifferentAbi(t *testing.T) {
	v := qbe.TypedValue{
		Value: qbe.TemporaryValue{Name: "name"},
		Type:  qbe.Byte,
	}

	if v.String() != "b %name" {
		t.Errorf("Expected b %%name, got %v", v.String())
	}

	if v.AbiString() != "w %name" {
		t.Errorf("Expected w %%name, got %v", v.AbiString())
	}
}
