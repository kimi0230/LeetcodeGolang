package cyclicrotation

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	arg2 int
	want []int
}{
	{
		[]int{3, 8, 9, 7, 6},
		3,
		[]int{9, 7, 6, 3, 8},
	},
	{
		[]int{0, 0, 0},
		1,
		[]int{0, 0, 0},
	},
	{
		[]int{1, 2, 3, 4},
		4,
		[]int{1, 2, 3, 4},
	},
	{
		[]int{-1, 2, 3, 4, 7, 10, -10},
		8,
		[]int{-10, -1, 2, 3, 4, 7, 10},
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
