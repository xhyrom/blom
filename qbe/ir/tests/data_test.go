package tests

import (
	"testing"

	"github.com/xhyrom/blom/qbe/ir"
)

func TestStringDataItem(t *testing.T) {
	item := ir.StringDataItem{Value: "hello"}
	expected := "\"hello\""
	if item.String() != expected {
		t.Errorf("expected %s, got %s", expected, item.String())
	}
}

func TestConstantDataItem(t *testing.T) {
	item := ir.ConstantDataItem{Value: 42}
	expected := "42"
	if item.String() != expected {
		t.Errorf("expected %s, got %s", expected, item.String())
	}
}

func TestDataString(t *testing.T) {
	align := uint64(8)
	data := ir.Data{
		Linkage: ir.NewLinkage(true),
		Name:    "myData",
		Align:   &align,
		Items: []ir.TypedDataItem{
			{Item: ir.StringDataItem{Value: "hello"}, Type: ir.NewPointer(ir.Char)},
			{Item: ir.ConstantDataItem{Value: 42}, Type: ir.Long},
		},
	}

	expected := "export data $myData = align 8 { l \"hello\", l 42 }"
	if data.String() != expected {
		t.Errorf("expected %s, got %s", expected, data.String())
	}
}

func TestDataStringNoAlign(t *testing.T) {
	data := ir.Data{
		Linkage: ir.NewLinkage(true),
		Name:    "myData",
		Align:   nil,
		Items: []ir.TypedDataItem{
			{Item: ir.StringDataItem{Value: "hello"}, Type: ir.NewPointer(ir.Char)},
			{Item: ir.ConstantDataItem{Value: 42}, Type: ir.Long},
		},
	}

	expected := "export data $myData = { l \"hello\", l 42 }"
	if data.String() != expected {
		t.Errorf("expected %s, got %s", expected, data.String())
	}
}

func TestDataStringNoLinkage(t *testing.T) {
	data := ir.Data{
		Linkage: ir.NewLinkage(false),
		Name:    "myData",
		Align:   nil,
		Items: []ir.TypedDataItem{
			{Item: ir.StringDataItem{Value: "hello"}, Type: ir.NewPointer(ir.Char)},
			{Item: ir.ConstantDataItem{Value: 42}, Type: ir.Long},
		},
	}

	expected := "data $myData = { l \"hello\", l 42 }"
	if data.String() != expected {
		t.Errorf("expected %s, got %s", expected, data.String())
	}
}
