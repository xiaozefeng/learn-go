package easy

import "testing"

func TestTowSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2,7,11,15},9, []int{0,1}},
		{[]int{2,7,11,15},13, []int{0,2}},
		{[]int{2,7,11,15},17, []int{0,3}},
	}

	for _, tt :=range tests{
		if actual := TowSum(tt.target,tt.nums); actual[0] != tt.expected[0] && actual[1] != tt.expected[1]{
			t.Errorf("expected: %+v, actual: %+v", tt.expected, actual)
		}
	}
}
