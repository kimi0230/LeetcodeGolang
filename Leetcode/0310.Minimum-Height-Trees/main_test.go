package minimumheighttrees

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 int
	arg2 [][]int
	want []int
}{
	{
		4,
		[][]int{{1, 0}, {1, 2}, {1, 3}},
		[]int{1},
	},
	{
		6,
		[][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
		[]int{3, 4},
	},
}

func TestFindMinHeightTrees(t *testing.T) {
	for _, tt := range tests {
		if got := FindMinHeightTrees(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}
