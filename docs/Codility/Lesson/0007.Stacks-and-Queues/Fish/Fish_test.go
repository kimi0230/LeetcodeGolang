package Fish

import "testing"

var tests = []struct {
	arg1 []int
	arg2 []int
	want int
}{
	{
		[]int{4, 3, 2, 1, 5},
		[]int{0, 1, 0, 0, 0},
		2,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
