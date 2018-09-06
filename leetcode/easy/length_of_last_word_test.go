package easy

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests:=[]struct{
		input string
		out int
	}{
		{"Hello World", 5},
		{"ab ", 2},
	}

	for _,tt := range tests{
		if actual := LengthOfLastWord(tt.input); actual != tt.out{
			t.Errorf("got %d, expected: %d", actual, tt.out)
		}
	}
}

