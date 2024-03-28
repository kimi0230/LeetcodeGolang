package PassingCars

import (
	"testing"
)

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{0, 1, 0, 1, 1},
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
