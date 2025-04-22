package tests

import (
	"testing"

	"github.com/xhyrom/blom/qbe/ir"
)

func TestComparisonTypeString(t *testing.T) {
	tests := []struct {
		name string
		c    ir.ComparisonType
		want string
	}{
		{
			name: "LessThan",
			c:    ir.LessThan,
			want: "lt",
		},
		{
			name: "LessThanOrEqual",
			c:    ir.LessThanOrEqual,
			want: "le",
		},
		{
			name: "GreaterThan",
			c:    ir.GreaterThan,
			want: "gt",
		},
		{
			name: "GreaterThanOrEqual",
			c:    ir.GreaterThanOrEqual,
			want: "ge",
		},
		{
			name: "Equal",
			c:    ir.Equal,
			want: "eq",
		},
		{
			name: "NotEqual",
			c:    ir.NotEqual,
			want: "ne",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("ComparisonType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnknownComparisonType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ComparisonType.String() did not panic")
		}
	}()

	_ = ir.ComparisonType(100).String()
}
