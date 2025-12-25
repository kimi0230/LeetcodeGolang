package MaxProfit

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{23171, 21011, 21123, 21366, 21013, 21367},
		356,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
