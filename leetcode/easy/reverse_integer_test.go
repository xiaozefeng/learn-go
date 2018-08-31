package easy

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct{ a, b int }{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
		{-2147483648, 0},
	}

	for _, tt := range tests{
		if actual := Reverse(tt.a); actual != tt.b{
			t.Errorf("expected: %d, actual: %d", tt.b, actual)
		}
	}


}
