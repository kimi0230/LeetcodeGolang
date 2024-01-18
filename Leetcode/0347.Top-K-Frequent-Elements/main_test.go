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

func TestTopKFrequentQuickSort(t *testing.T) {
	for _, tt := range tests {
		if got := TopKFrequentQuickSort(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
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

func BenchmarkTopKFrequentQuickSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopKFrequentQuickSort(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0347.Top-K-Frequent-Elements -bench=.
go test -benchmem -run=none LeetcodeGolang/Leetcode/0347.Top-K-Frequent-Elements -bench=.            100%    14:39:46 
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0347.Top-K-Frequent-Elements
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkTopKFrequent-8                  1875834               648.6 ns/op           200 B/op         11 allocs/op
BenchmarkTopKFrequentQuickSort-8         1830498               704.2 ns/op           312 B/op         11 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0347.Top-K-Frequent-Elements    3.831s
*/
