package tests

import (
	"blom/qbe"
	"testing"
)

func TestStringDataItem(t *testing.T) {
	item := qbe.StringDataItem{Value: "hello"}
	expected := "\"hello\""
	if item.String() != expected {
		t.Errorf("expected %s, got %s", expected, item.String())
	}
}

func TestConstantDataItem(t *testing.T) {
	item := qbe.ConstantDataItem{Value: 42}
	expected := "42"
	if item.String() != expected {
		t.Errorf("expected %s, got %s", expected, item.String())
	}
}

func TestDataString(t *testing.T) {
	align := uint64(8)
	data := qbe.Data{
		Linkage: qbe.NewLinkage(true),
		Name:    "myData",
		Align:   &align,
		Items: []qbe.TypedDataItem{
			{Item: qbe.StringDataItem{Value: "hello"}, Type: qbe.String},
			{Item: qbe.ConstantDataItem{Value: 42}, Type: qbe.Long},
		},
	}

	expected := "export data $myData = align 8 { l \"hello\", l 42 }"
	if data.String() != expected {
		t.Errorf("expected %s, got %s", expected, data.String())
	}
}

func TestDataStringNoAlign(t *testing.T) {
	data := qbe.Data{
		Linkage: qbe.NewLinkage(true),
		Name:    "myData",
		Align:   nil,
		Items: []qbe.TypedDataItem{
			{Item: qbe.StringDataItem{Value: "hello"}, Type: qbe.String},
			{Item: qbe.ConstantDataItem{Value: 42}, Type: qbe.Long},
		},
	}

	expected := "export data $myData = { l \"hello\", l 42 }"
	if data.String() != expected {
		t.Errorf("expected %s, got %s", expected, data.String())
	}
}

func TestDataStringNoLinkage(t *testing.T) {
	data := qbe.Data{
		Linkage: qbe.NewLinkage(false),
		Name:    "myData",
		Align:   nil,
		Items: []qbe.TypedDataItem{
			{Item: qbe.StringDataItem{Value: "hello"}, Type: qbe.String},
			{Item: qbe.ConstantDataItem{Value: 42}, Type: qbe.Long},
		},
	}

	expected := "data $myData = { l \"hello\", l 42 }"
	if data.String() != expected {
		t.Errorf("expected %s, got %s", expected, data.String())
	}
}
