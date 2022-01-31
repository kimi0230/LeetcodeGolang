package longestincreasingsubsequence

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{10, 9, 2, 5, 3, 7, 101, 18},
		4,
	},
	{
		[]int{0, 1, 0, 3, 2, 3},
		4,
	},
	{
		[]int{7, 7, 7, 7, 7, 7, 7},
		1,
	},
}

func TestLengthOfLIS(t *testing.T) {
	for _, tt := range tests {
		if got := LengthOfLIS(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLengthOfLIS2(t *testing.T) {
	for _, tt := range tests {
		if got := LengthOfLIS2(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
