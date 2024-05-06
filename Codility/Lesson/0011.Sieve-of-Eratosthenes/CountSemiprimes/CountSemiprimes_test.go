package countsemiprimes

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
		26,
		[]int{1, 4, 16},
		[]int{26, 10, 20},
		[]int{10, 4, 0},
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2, tt.arg3); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, tt := range tests {
		if got := Solution2(tt.arg1, tt.arg2, tt.arg3); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution(tests[0].arg1, tests[0].arg2, tests[0].arg3)
	}
}

func BenchmarkSolution2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution(tests[0].arg1, tests[0].arg2, tests[0].arg3)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Codility/Lesson/0011.Sieve-of-Eratosthenes/CountSemiprimes -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Codility/Lesson/0011.Sieve-of-Eratosthenes/CountSemiprimes
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkSolution-8       316700              3294 ns/op            1907 B/op         15 allocs/op
BenchmarkSolution2-8      361731              3326 ns/op            1906 B/op         15 allocs/op
PASS
ok      LeetcodeGolang/Codility/Lesson/0011.Sieve-of-Eratosthenes/CountSemiprimes       2.577s
*/
