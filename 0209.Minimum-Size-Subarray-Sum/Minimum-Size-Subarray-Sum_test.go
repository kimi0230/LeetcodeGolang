package minimumsizesubarraysum

import (
	"testing"
)

var tests = []struct {
	arg1 int
	arg2 []int
	want int
}{
	{
		7,
		[]int{2, 3, 1, 2, 4, 3},
		2,
	},
	{
		4,
		[]int{1, 4, 4},
		1,
	},
	{
		11,
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		0,
	},
}

func TestMinSubArrayLenBurst(t *testing.T) {
	for _, tt := range tests {
		if got := MinSubArrayLenBurst(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}
