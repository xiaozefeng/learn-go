package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{2, 2, 4},
		{12, 21, 33},
	}

	for _, tt := range tests {
		if actual := add(tt.a, tt.b); actual != tt.c {
			t.Errorf("add(%d, %d); got %d; expected %d ",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
