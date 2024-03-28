package commonprimedivisors

import "testing"

var tests = []struct {
	arg1 []int
	arg2 []int
	want int
}{
	{
		[]int{15, 10, 3},
		[]int{75, 30, 5},
		1,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
