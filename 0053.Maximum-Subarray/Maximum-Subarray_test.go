package maximumsubarray

import "testing"

func TestMaxSubArray1(t *testing.T) {
	tests := []struct {
		arg1 []int
		want int
	}{
		{
			arg1: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
	}
	for _, tt := range tests {
		if got := MaxSubArray1(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
