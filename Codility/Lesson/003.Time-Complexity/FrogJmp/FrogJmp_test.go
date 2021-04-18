package frogjump

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 int
	arg2 int
	arg3 int
	want int
}{
	{
		10, 80, 30, 3,
	},
	{
		50, 40, 30, 0,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2, tt.arg3); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
