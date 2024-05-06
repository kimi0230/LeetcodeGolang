package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {

	tests := []struct {
		arg1 []int
		arg2 int
		arg3 []int
		arg4 int
		want []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, tt := range tests {
		Merge(tt.arg1, tt.arg2, tt.arg3, tt.arg4)
		if !reflect.DeepEqual(tt.arg1, tt.want) {
			t.Errorf("got = %v, want = %v", tt.arg1, tt.want)
		}
	}

}
