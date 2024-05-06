package Triangle

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{10, 2, 5, 1, 8, 20},
		1,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
