package tests

import (
	"blom/qbe"
	"testing"
)

func TestComparisonTypeString(t *testing.T) {
	tests := []struct {
		name string
		c    qbe.ComparisonType
		want string
	}{
		{
			name: "LessThan",
			c:    qbe.LessThan,
			want: "lt",
		},
		{
			name: "LessThanOrEqual",
			c:    qbe.LessThanOrEqual,
			want: "le",
		},
		{
			name: "GreaterThan",
			c:    qbe.GreaterThan,
			want: "gt",
		},
		{
			name: "GreaterThanOrEqual",
			c:    qbe.GreaterThanOrEqual,
			want: "ge",
		},
		{
			name: "Equal",
			c:    qbe.Equal,
			want: "eq",
		},
		{
			name: "NotEqual",
			c:    qbe.NotEqual,
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

	_ = qbe.ComparisonType(100).String()
}
