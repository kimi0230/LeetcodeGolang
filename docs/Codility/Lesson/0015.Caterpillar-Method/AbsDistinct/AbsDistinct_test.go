package absdistinct

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{-5, -3, -1, 0, 3, 6},
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
