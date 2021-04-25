package Distinct

import (
	"testing"
)

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{2, 1, 1, 2, 3, 1},
		3,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolutionSet(t *testing.T) {
	for _, tt := range tests {
		if got := SolutionSet(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution(tests[0].arg1)
	}
}

func BenchmarkSolutionSet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SolutionSet(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Codility/Lesson/0006.Sorting/Distinct -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Codility/Lesson/0006.Sorting/Distinct
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkSolution-4             15742278                70.30 ns/op           24 B/op          1 allocs/op
BenchmarkSolutionSet-4          14324376                81.14 ns/op            0 B/op          0 allocs/op
*/
