package PermCheck

import (
	"testing"
)

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{4, 1, 3, 2},
		1,
	},
	{
		[]int{4, 3, 2},
		0,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
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
