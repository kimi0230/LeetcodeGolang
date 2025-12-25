package MaxCounters

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 int
	arg2 []int
	want []int
}{
	{
		5,
		[]int{3, 4, 4, 6, 1, 4, 4},
		[]int{3, 2, 2, 4, 2},
	},
	{
		5,
		[]int{3, 4, 4, 6, 1, 4, 4},
		[]int{3, 2, 2, 4, 2},
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
