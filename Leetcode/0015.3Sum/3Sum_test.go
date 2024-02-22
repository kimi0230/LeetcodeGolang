package threesum

import (
	"reflect"
	"sort"
	"testing"
)

var tests = []struct {
	arg1 []int
	want [][]int
}{
	{
		[]int{0, 0, 0},
		[][]int{{0, 0, 0}},
	},
	{
		[]int{0, 0, 0, 0},
		[][]int{{0, 0, 0}},
	},
	{
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		[]int{},
		[][]int{},
	},
	{
		[]int{0},
		[][]int{},
	},
}

func TestThreeSumBurst(t *testing.T) {
	for _, tt := range tests {
		if got := ThreeSumBurst(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func TestThreeSumDoublePoint(t *testing.T) {
	for _, tt := range tests {
		if got := ThreeSumDoublePoint(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

// 判断切片在词典序上的大小关系
func lexicographicLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return len(a) < len(b)
}

// 对切片的切片进行排序
func sortSliceOfSlices(s [][]int) [][]int {
	sort.Slice(s, func(i, j int) bool {
		return lexicographicLess(s[i], s[j])
	})
	return s
}

func TestThreeSumHashTable(t *testing.T) {
	for _, tt := range tests {
		if got := ThreeSumHashTable(tt.arg1); !reflect.DeepEqual(sortSliceOfSlices(got), tt.want) {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func TestThreeSumTwoPointer(t *testing.T) {
	for _, tt := range tests {
		if got := ThreeSumTwoPointer(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func BenchmarkThreeSumBurst(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ThreeSumBurst(tests[0].arg1)
	}
}

func BenchmarkThreeSumDoublePoint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ThreeSumDoublePoint(tests[0].arg1)
	}
}

func BenchmarkThreeSumHashTable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ThreeSumHashTable(tests[0].arg1)
	}
}

func BenchmarkThreeSumTwoPointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ThreeSumTwoPointer(tests[0].arg1)
	}
}

/*
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0015.3Sum
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkThreeSumBurst-4                 9838000               121.4 ns/op            48 B/op          2 allocs/op
BenchmarkThreeSumDoublePoint-4           9069201               112.8 ns/op            48 B/op          2 allocs/op
BenchmarkThreeSumHashTable-4             7935907               147.1 ns/op            48 B/op          2 allocs/op
BenchmarkThreeSumTwoPointer-4           10888315               103.5 ns/op            48 B/op          2 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0015.3Sum       5.055s
*/
