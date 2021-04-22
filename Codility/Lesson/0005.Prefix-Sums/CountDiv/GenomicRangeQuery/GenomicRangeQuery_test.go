package GenomicRangeQuery

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 string
	arg2 []int
	arg3 []int
	want []int
}{
	{
		"CAGCCTA",
		[]int{2, 5, 0},
		[]int{4, 5, 6},
		[]int{2, 4, 1},
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2, tt.arg3); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
