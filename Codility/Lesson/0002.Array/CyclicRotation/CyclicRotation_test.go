package cyclicrotation

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
		[]int{3, 8, 9, 7, 6},
		3,
		[]int{9, 7, 6, 3, 8},
	},
	{
		[]int{0, 0, 0},
		1,
		[]int{0, 0, 0},
	},
	{
		[]int{1, 2, 3, 4},
		4,
		[]int{1, 2, 3, 4},
	},
	{
		[]int{-1, 2, 3, 4, 7, 10, -10},
		8,
		[]int{-10, -1, 2, 3, 4, 7, 10},
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution(tests[3].arg1, tests[3].arg2)
	}
}
func BenchmarkSolution2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution2(tests[3].arg1, tests[3].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Codility/Lesson/0002.Array/CyclicRotation -bench=.
BenchmarkSolution-4     19386909                56.24 ns/op           64 B/op          1 allocs/op
BenchmarkSolution2-4    20282828                54.88 ns/op           64 B/op          1 allocs/op
*/
