package NumberOfDiscIntersections

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{1, 5, 2, 1, 4, 0},
		11,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolutionDirect(t *testing.T) {
	for _, tt := range tests {
		if got := SolutionDirect(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
