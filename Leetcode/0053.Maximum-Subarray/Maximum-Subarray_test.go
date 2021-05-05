package maximumsubarray

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		arg1: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		want: 6,
	},
}

func TestMaxSubArrayDP(t *testing.T) {
	for _, tt := range tests {
		if got := MaxSubArrayDP(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestMaxSubArray1(t *testing.T) {

	for _, tt := range tests {
		if got := MaxSubArray1(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

var benchArg1 = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

func BenchmarkMaxSubArrayDP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxSubArrayDP(benchArg1)
	}
}
func BenchmarkMaxSubArray1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxSubArray1(benchArg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0053.Maximum-Subarray -bench=.
BenchmarkMaxSubArrayDP-8        22077684                48.19 ns/op           80 B/op          1 allocs/op
BenchmarkMaxSubArray1-8         94770036                10.74 ns/op            0 B/op          0 allocs/op
*/
