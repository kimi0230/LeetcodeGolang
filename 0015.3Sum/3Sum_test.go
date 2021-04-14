package threesum

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	want [][]int
}{
	{
		[]int{0, 0, 0},
		[][]int{{0, 0, 0}},
	},
	{
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		[]int{},
		[][]int{},
	},
	{
		[]int{0},
		[][]int{},
	},
}

func TestThreeSumBurst(t *testing.T) {
	for _, tt := range tests {
		if got := ThreeSumBurst(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}
