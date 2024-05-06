package MaxDoubleSliceSum

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{3, 2, 6, -1, 4, 5, -1, 2},
		17,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
