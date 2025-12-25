package Brackets

import "testing"

var tests = []struct {
	arg1 string
	want int
}{
	{
		"{[()()]}",
		1,
	},
	{
		"([)()]",
		0,
	},
	{
		")(",
		0,
	},
	{
		"",
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
