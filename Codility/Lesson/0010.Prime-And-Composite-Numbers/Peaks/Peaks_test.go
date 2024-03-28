package peaks

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
		3,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
