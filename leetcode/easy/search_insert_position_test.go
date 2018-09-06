package easy

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		input []int
		target  int
		expected int
	}{
		{[]int{1,3,5,6}, 5 ,2},
		{[]int{1,3,5,6}, 2  ,1},
		{[]int{1,3,5,6}, 7, 4 },
		{[]int{1,3,5,6}, 0 ,0},
	}



	for _, tt := range tests{
		if actual := SearchInsert(tt.input,tt.target); actual != tt.expected{
			t.Errorf("got: %d, expected %d", actual, tt.expected)
		}
	}
}
