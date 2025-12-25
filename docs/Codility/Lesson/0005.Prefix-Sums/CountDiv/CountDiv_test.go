package CountDiv

import (
	"testing"
)

var tests = []struct {
	arg1 int
	arg2 int
	arg3 int
	want int
}{
	{
		6, 11, 2, 3,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2, tt.arg3); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolutionBurst(t *testing.T) {
	for _, tt := range tests {
		if got := SolutionBurst(tt.arg1, tt.arg2, tt.arg3); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
