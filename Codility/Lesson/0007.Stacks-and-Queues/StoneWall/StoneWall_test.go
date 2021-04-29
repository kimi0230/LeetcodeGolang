package StoneWall

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{8, 8, 5, 7, 9, 8, 7, 4, 8},
		7,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
