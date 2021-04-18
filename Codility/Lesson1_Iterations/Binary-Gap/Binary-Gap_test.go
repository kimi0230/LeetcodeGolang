package binarygap

import "testing"

var tests = []struct {
	arg1 int
	want int
}{
	{
		141, // 1000 1101
		3,
	},
	{
		1041, // 10100 0001 0001
		5,
	},
	{
		32, // 10100 0001 0001
		0,
	},
	{
		2147483647, // 2^31-1: 0111 1111 1111 1111 1111 1111 1111 1111
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
