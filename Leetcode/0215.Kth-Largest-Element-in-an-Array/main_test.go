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
	{
		[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		4,
		4,
	},
}

func TestFindKthLargestHeap(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := FindKthLargestHeap(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestFindKthLargestSort(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := FindKthLargestSort(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkFindKthLargestHeap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindKthLargestHeap(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkFindKthLargestSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindKthLargestSort(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0215.Kth-Largest-Element-in-an-Array -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0215.Kth-Largest-Element-in-an-Array
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkFindKthLargestHeap-8            6112726               184.9 ns/op            48 B/op          3 allocs/op
BenchmarkFindKthLargestSort-8           18023403                60.38 ns/op           24 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0215.Kth-Largest-Element-in-an-Array    3.383s
*/
