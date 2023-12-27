package findtargetlastindex

import "testing"

var tests = []struct {
	arg1 []int
	arg2 int
	want int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 6, 6, 9},
		6,
		7,
	},
	{
		[]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 8, 9, 9, 9, 9},
		6,
		12,
	},
	{
		[]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 8, 9, 9, 9, 9},
		5,
		6,
	},
	{
		[]int{5, 7, 7, 8, 8, 10},
		8,
		4,
	},
	{
		[]int{5, 7, 7, 8, 8, 10},
		6,
		-1,
	},
	{
		[]int{1, 2, 3, 4, 5},
		3,
		2,
	},
	{
		[]int{1, 2, 3, 4, 5},
		5,
		4,
	},
	{[]int{1, 2, 3, 4, 5},
		0,
		-1,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolutionRecursive(t *testing.T) {
	for _, tt := range tests {
		// start := 0
		// end := len(tt.arg1) - 1
		if got := SolutionRecursive(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}

		// SolutionRecursive(tt.arg1, start, end, tt.arg2)
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
func BenchmarkSolutionRecursive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SolutionRecursive(tt.arg1, tt.arg2)
		}
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Algorithms/Find_Target_Last_Index -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Algorithms/Find_Target_Last_Index
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkSolution-8             54216429                24.75 ns/op            0 B/op          0 allocs/op
BenchmarkSolutionRecursive-8    40756744                27.75 ns/op            0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Algorithms/Find_Target_Last_Index        2.537s
*/
