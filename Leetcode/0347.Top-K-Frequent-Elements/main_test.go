package topkfrequentelements

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	arg2 int
	want []int
}{
	{
		[]int{1, 1, 1, 2, 2, 3},
		2,
		[]int{1, 2},
	},
	{
		[]int{1},
		1,
		[]int{1},
	},
}

func TestTopKFrequent(t *testing.T) {
	for _, tt := range tests {
		if got := TopKFrequent(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopKFrequent(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
