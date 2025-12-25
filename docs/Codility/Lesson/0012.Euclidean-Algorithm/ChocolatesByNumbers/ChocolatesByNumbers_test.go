package chocolatesbynumbers

import "testing"

var tests = []struct {
	arg1 int
	arg2 int
	want int
}{
	{
		10,
		4,
		5,
	},
	{
		947853,
		4453,
		947853,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolutionBurst(t *testing.T) {
	for _, tt := range tests {
		if got := SolutionBurst(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			Solution(tt.arg1, tt.arg2)
		}
	}
}

func BenchmarkSolutionBurst(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SolutionBurst(tt.arg1, tt.arg2)
		}
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Codility/Lesson/0012.Euclidean-Algorithm/ChocolatesByNumbers -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Codility/Lesson/0012.Euclidean-Algorithm/ChocolatesByNumbers
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkSolution-4             12411740               112.2 ns/op             0 B/op          0 allocs/op
BenchmarkSolutionBurst-4               3         354775339 ns/op        49782528 B/op      38396 allocs/op
PASS
ok      LeetcodeGolang/Codility/Lesson/0012.Euclidean-Algorithm/ChocolatesByNumbers     4.425s
*/
