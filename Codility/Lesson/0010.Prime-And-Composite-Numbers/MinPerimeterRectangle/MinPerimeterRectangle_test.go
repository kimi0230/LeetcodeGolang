package minperimeterrectangle

import "testing"

var tests = []struct {
	arg1 int
	want int
}{
	{
		30,
		22,
	},
	{
		1,
		4,
	},
	{
		36,
		24,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolution1(t *testing.T) {
	for _, tt := range tests {
		if got := Solution1(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, tt := range tests {
		if got := Solution2(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
