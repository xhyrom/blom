package qbe

import (
	"fmt"
	"strings"
)

type Data struct {
	Linkage Linkage
	Name    string
	Align   *uint64
	Items   []TypedDataItem
}

func (d Data) String() string {
	result := fmt.Sprintf("%sdata %s = ", d.Linkage, d.Name)

	if d.Align != nil {
		result += fmt.Sprintf("align %d ", *d.Align)
	}

	var parts []string
	for _, item := range d.Items {
		parts = append(parts, fmt.Sprintf("%s %s", item.Type, item.Item))
	}

	result += fmt.Sprintf("{ %s }", strings.Join(parts, ", "))

	return result
}

// Data represents a data section.
type DataItemType int

const (
	StringDataItemType DataItemType = iota
	ConstantDataItemType
)

type DataItem interface {
	DataType() DataItemType
	String() string
}

// StringDataItem represents a string data item. ("value")
type StringDataItem struct {
	Value string
}

func (i StringDataItem) DataType() DataItemType {
	return StringDataItemType
}

func (i StringDataItem) String() string {
	return fmt.Sprintf("\"%s\"", i.Value)
}

// ConstantDataItem represents a constant data item. (value)
type ConstantDataItem struct {
	Value int64
}

func (i ConstantDataItem) DataType() DataItemType {
	return ConstantDataItemType
}

func (i ConstantDataItem) String() string {
	return fmt.Sprintf("%d", i.Value)
}

// TypedDataItem represents a typed data item. (type value)
type TypedDataItem struct {
	Item DataItem
	Type Type
}
