package MinAvgTwoSlice

import (
	"testing"
)

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{4, 2, 2, 5, 1, 5, 8},
		1,
	},
	{
		[]int{1, 3, 6, 4, 1, 2},
		4,
	},
	{
		[]int{10, 10, -1, 2, 4, -1, 2, -1},
		5,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
