package tests

import (
	"testing"

	"github.com/xhyrom/blom/qbe/ir"
)

func TestTypeToString(t *testing.T) {
	tests := []struct {
		input    ir.Type
		expected string
	}{
		{ir.Byte, "b"},
		{ir.UnsignedByte, "ub"},
		{ir.Halfword, "h"},
		{ir.UnsignedHalfword, "uh"},
		{ir.Word, "w"},
		{ir.UnsignedWord, "uw"},
		{ir.Long, "l"},
		{ir.UnsignedLong, "ul"},
		{ir.Single, "s"},
		{ir.Double, "d"},
		{ir.Char, "b"},
		{ir.Boolean, "w"},
		{ir.Pointer, "l"},
		{ir.Void, "w"},
		{ir.Null, ""},
	}

	for _, test := range tests {
		if result := test.input.String(); result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestTypeIsNumeric(t *testing.T) {
	tests := []struct {
		input    ir.Type
		expected bool
	}{
		{ir.Byte, true},
		{ir.UnsignedByte, true},
		{ir.Halfword, true},
		{ir.UnsignedHalfword, true},
		{ir.Word, true},
		{ir.UnsignedWord, true},
		{ir.Long, true},
		{ir.UnsignedLong, true},
		{ir.Single, true},
		{ir.Double, true},
		{ir.Char, false},
		{ir.Boolean, false},
		{ir.Pointer, false},
		{ir.Void, false},
		{ir.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsNumeric(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIsInteger(t *testing.T) {
	tests := []struct {
		input    ir.Type
		expected bool
	}{
		{ir.Byte, true},
		{ir.UnsignedByte, true},
		{ir.Halfword, true},
		{ir.UnsignedHalfword, true},
		{ir.Word, true},
		{ir.UnsignedWord, true},
		{ir.Long, true},
		{ir.UnsignedLong, true},
		{ir.Single, false},
		{ir.Double, false},
		{ir.Char, false},
		{ir.Boolean, false},
		{ir.Pointer, false},
		{ir.Void, false},
		{ir.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsInteger(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIsFloatingPoint(t *testing.T) {
	tests := []struct {
		input    ir.Type
		expected bool
	}{
		{ir.Byte, false},
		{ir.UnsignedByte, false},
		{ir.Halfword, false},
		{ir.UnsignedHalfword, false},
		{ir.Word, false},
		{ir.UnsignedWord, false},
		{ir.Long, false},
		{ir.UnsignedLong, false},
		{ir.Single, true},
		{ir.Double, true},
		{ir.Char, false},
		{ir.Boolean, false},
		{ir.Pointer, false},
		{ir.Void, false},
		{ir.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsFloatingPoint(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIsSigned(t *testing.T) {
	tests := []struct {
		input    ir.Type
		expected bool
	}{
		{ir.Byte, true},
		{ir.UnsignedByte, false},
		{ir.Halfword, true},
		{ir.UnsignedHalfword, false},
		{ir.Word, true},
		{ir.UnsignedWord, false},
		{ir.Long, true},
		{ir.UnsignedLong, false},
		{ir.Single, false},
		{ir.Double, false},
		{ir.Char, false},
		{ir.Boolean, false},
		{ir.Pointer, false},
		{ir.Void, false},
		{ir.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsSigned(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIsUnsigned(t *testing.T) {
	tests := []struct {
		input    ir.Type
		expected bool
	}{
		{ir.Byte, false},
		{ir.UnsignedByte, true},
		{ir.Halfword, false},
		{ir.UnsignedHalfword, true},
		{ir.Word, false},
		{ir.UnsignedWord, true},
		{ir.Long, false},
		{ir.UnsignedLong, true},
		{ir.Single, false},
		{ir.Double, false},
		{ir.Char, false},
		{ir.Boolean, false},
		{ir.Pointer, false},
		{ir.Void, false},
		{ir.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsUnsigned(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIntoAbi(t *testing.T) {
	tests := []struct {
		input    ir.Type
		expected ir.Type
	}{
		{ir.Byte, ir.Word},
		{ir.Char, ir.Word},
		{ir.UnsignedByte, ir.Word},
		{ir.Halfword, ir.Word},
		{ir.UnsignedHalfword, ir.Word},
		{ir.UnsignedWord, ir.Word},
		{ir.Word, ir.Word},
		{ir.Long, ir.Long},
		{ir.UnsignedLong, ir.UnsignedLong},
		{ir.Single, ir.Single},
		{ir.Double, ir.Double},
		{ir.Boolean, ir.Boolean},
		{ir.Pointer, ir.Pointer},
		{ir.Void, ir.Void},
		{ir.Null, ir.Null},
	}

	for _, test := range tests {
		if result := test.input.IntoAbi(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeDefinitionString(t *testing.T) {
	align := uint64(8)
	tests := []struct {
		input    ir.TypeDefinition
		expected string
	}{
		{
			input: ir.TypeDefinition{
				Name:  "test",
				Align: &align,
				Items: []ir.TypedTypeDefinitionItem{
					{Count: 1, Type: ir.Word},
					{Count: 2, Type: ir.Byte},
				},
			},
			expected: "type :test = align 8 { w, b 2 }",
		},
		{
			input: ir.TypeDefinition{
				Name: "test2",
				Items: []ir.TypedTypeDefinitionItem{
					{Count: 1, Type: ir.Word},
					{Count: 1, Type: ir.Byte},
				},
			},
			expected: "type :test2 = { w, b }",
		},
	}

	for _, test := range tests {
		if result := test.input.String(); result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
