package permutations

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	want [][]int
}{
	{
		arg1: []int{1, 2, 3},
		want: [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		},
	},
}

func TestPermutations(t *testing.T) {
	for _, tt := range tests {
		if got := Permute(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
