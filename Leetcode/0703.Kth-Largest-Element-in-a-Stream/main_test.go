package kthlargestelementinastream

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
		3,
		[]int{4, 5, 8, 2},
		[]int{3, 5, 10, 9, 4},
		[]int{0, 4, 5, 5, 8, 8},
	},
}

func TestKthLargestStream(t *testing.T) {
	for _, tt := range tests {
		if got := KthLargestStream(tt.arg1, tt.arg2, tt.arg3); !reflect.DeepEqual(got, tt.want) {
			// if got := KthLargestStream(tt.arg1, tt.arg2, tt.arg3); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkKthLargestStream(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KthLargestStream(tests[0].arg1, tests[0].arg2, tests[0].arg3)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
