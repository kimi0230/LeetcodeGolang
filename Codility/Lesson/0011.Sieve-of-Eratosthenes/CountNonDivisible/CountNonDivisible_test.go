package countnondivisible

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	want []int
}{
	{
		[]int{3, 1, 2, 3, 6},
		[]int{2, 4, 3, 2, 0},
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, tt := range tests {
		if got := Solution2(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolution1(t *testing.T) {
	for _, tt := range tests {
		if got := Solution1(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
