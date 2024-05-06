package FrogRiverOne

import "testing"

var tests = []struct {
	arg1 int
	arg2 []int
	want int
}{
	{
		5,
		[]int{1, 3, 1, 4, 2, 3, 5, 4},
		6,
	},
	{
		2,
		[]int{2, 2, 2, 2, 2},
		-1,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
