package kthlargestelementinanarray

import "testing"

var tests = []struct {
	arg1 []int
	arg2 int
	want int
}{
	{
		[]int{3, 2, 1, 5, 6, 4},
		2,
		5,
	},
}

func TestFindKthLargest(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := FindKthLargest(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindKthLargest(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
