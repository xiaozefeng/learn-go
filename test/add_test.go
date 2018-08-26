package main

import (
	"testing"
)

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

func BenchmarkAdd(b *testing.B) {
	a, c := 1, 2
	ans := 3
	actual := add(a, c)
	for i := 0; i < b.N; i++ {
		if ans != actual {
			b.Errorf("got %d for input %d expected %d", actual, ans, actual)
		}
	}
}

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcab", 3},
		{"pwwkew", 3},
		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},

		// Chinese support
		{"这里是", 3},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; experted %d",
				actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s experted %d", actual, s, ans)
		}
	}
}
