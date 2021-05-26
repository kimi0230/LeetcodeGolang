package countsemiprimes

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 int
	arg2 []int
	arg3 []int
	want []int
}{
	{
		26,
		[]int{1, 4, 16},
		[]int{26, 10, 20},
		[]int{10, 4, 0},
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2, tt.arg3); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, tt := range tests {
		if got := Solution2(tt.arg1, tt.arg2, tt.arg3); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
