package spiralmatrixii

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 int
	want [][]int
}{
	{
		2,
		[][]int{{1, 2}, {4, 3}},
	},
	{
		3,
		[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
	},
	{
		4,
		[][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
	},
}

func TestGenerateMatrix(t *testing.T) {
	for _, tt := range tests {
		if got := GenerateMatrix(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
