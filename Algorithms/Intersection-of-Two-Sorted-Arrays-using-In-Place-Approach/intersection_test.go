package intersection

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	var tests = []struct {
		arg1 []int
		arg2 []int
		want []int
	}{
		{
			arg1: []int{1, 3, 4, 6, 7},
			arg2: []int{2, 4, 6, 8, 9},
			want: []int{4, 6},
		},
	}

	for _, tt := range tests {
		if got := FindIntersection(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
