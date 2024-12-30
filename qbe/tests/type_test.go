package tests

import (
	"blom/qbe"
	"testing"
)

func TestTypeToString(t *testing.T) {
	tests := []struct {
		input    qbe.Type
		expected string
	}{
		{qbe.Byte, "b"},
		{qbe.UnsignedByte, "ub"},
		{qbe.Halfword, "h"},
		{qbe.UnsignedHalfword, "uh"},
		{qbe.Word, "w"},
		{qbe.UnsignedWord, "uw"},
		{qbe.Long, "l"},
		{qbe.UnsignedLong, "ul"},
		{qbe.Single, "s"},
		{qbe.Double, "d"},
		{qbe.Char, "b"},
		{qbe.Boolean, "w"},
		{qbe.String, "l"},
		{qbe.Void, "w"},
		{qbe.Null, ""},
	}

	for _, test := range tests {
		if result := test.input.String(); result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestTypeIsNumeric(t *testing.T) {
	tests := []struct {
		input    qbe.Type
		expected bool
	}{
		{qbe.Byte, true},
		{qbe.UnsignedByte, true},
		{qbe.Halfword, true},
		{qbe.UnsignedHalfword, true},
		{qbe.Word, true},
		{qbe.UnsignedWord, true},
		{qbe.Long, true},
		{qbe.UnsignedLong, true},
		{qbe.Single, true},
		{qbe.Double, true},
		{qbe.Char, false},
		{qbe.Boolean, false},
		{qbe.String, false},
		{qbe.Void, false},
		{qbe.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsNumeric(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIsInteger(t *testing.T) {
	tests := []struct {
		input    qbe.Type
		expected bool
	}{
		{qbe.Byte, true},
		{qbe.UnsignedByte, true},
		{qbe.Halfword, true},
		{qbe.UnsignedHalfword, true},
		{qbe.Word, true},
		{qbe.UnsignedWord, true},
		{qbe.Long, true},
		{qbe.UnsignedLong, true},
		{qbe.Single, false},
		{qbe.Double, false},
		{qbe.Char, false},
		{qbe.Boolean, false},
		{qbe.String, false},
		{qbe.Void, false},
		{qbe.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsInteger(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIsFloatingPoint(t *testing.T) {
	tests := []struct {
		input    qbe.Type
		expected bool
	}{
		{qbe.Byte, false},
		{qbe.UnsignedByte, false},
		{qbe.Halfword, false},
		{qbe.UnsignedHalfword, false},
		{qbe.Word, false},
		{qbe.UnsignedWord, false},
		{qbe.Long, false},
		{qbe.UnsignedLong, false},
		{qbe.Single, true},
		{qbe.Double, true},
		{qbe.Char, false},
		{qbe.Boolean, false},
		{qbe.String, false},
		{qbe.Void, false},
		{qbe.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsFloatingPoint(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIsSigned(t *testing.T) {
	tests := []struct {
		input    qbe.Type
		expected bool
	}{
		{qbe.Byte, true},
		{qbe.UnsignedByte, false},
		{qbe.Halfword, true},
		{qbe.UnsignedHalfword, false},
		{qbe.Word, true},
		{qbe.UnsignedWord, false},
		{qbe.Long, true},
		{qbe.UnsignedLong, false},
		{qbe.Single, false},
		{qbe.Double, false},
		{qbe.Char, false},
		{qbe.Boolean, false},
		{qbe.String, false},
		{qbe.Void, false},
		{qbe.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsSigned(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIsUnsigned(t *testing.T) {
	tests := []struct {
		input    qbe.Type
		expected bool
	}{
		{qbe.Byte, false},
		{qbe.UnsignedByte, true},
		{qbe.Halfword, false},
		{qbe.UnsignedHalfword, true},
		{qbe.Word, false},
		{qbe.UnsignedWord, true},
		{qbe.Long, false},
		{qbe.UnsignedLong, true},
		{qbe.Single, false},
		{qbe.Double, false},
		{qbe.Char, false},
		{qbe.Boolean, false},
		{qbe.String, false},
		{qbe.Void, false},
		{qbe.Null, false},
	}

	for _, test := range tests {
		if result := test.input.IsUnsigned(); result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestTypeIntoAbi(t *testing.T) {
	tests := []struct {
		input    qbe.Type
		expected qbe.Type
	}{
		{qbe.Byte, qbe.Word},
		{qbe.Char, qbe.Word},
		{qbe.UnsignedByte, qbe.Word},
		{qbe.Halfword, qbe.Word},
		{qbe.UnsignedHalfword, qbe.Word},
		{qbe.UnsignedWord, qbe.Word},
		{qbe.Word, qbe.Word},
		{qbe.Long, qbe.Long},
		{qbe.UnsignedLong, qbe.UnsignedLong},
		{qbe.Single, qbe.Single},
		{qbe.Double, qbe.Double},
		{qbe.Boolean, qbe.Boolean},
		{qbe.String, qbe.String},
		{qbe.Void, qbe.Void},
		{qbe.Null, qbe.Null},
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
		input    qbe.TypeDefinition
		expected string
	}{
		{
			input: qbe.TypeDefinition{
				Name:  "test",
				Align: &align,
				Items: []qbe.TypedTypeDefinitionItem{
					{Count: 1, Type: qbe.Word},
					{Count: 2, Type: qbe.Byte},
				},
			},
			expected: "type :test = align 8 { w, b 2 }",
		},
		{
			input: qbe.TypeDefinition{
				Name: "test2",
				Items: []qbe.TypedTypeDefinitionItem{
					{Count: 1, Type: qbe.Word},
					{Count: 1, Type: qbe.Byte},
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
