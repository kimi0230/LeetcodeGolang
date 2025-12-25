package sortcolors

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	want []int
}{
	{
		[]int{2, 0, 2, 1, 1, 0},
		[]int{0, 0, 1, 1, 2, 2},
	},
	{
		[]int{2, 0, 1},
		[]int{0, 1, 2},
	},
	{
		[]int{0},
		[]int{0},
	},
	{
		[]int{1},
		[]int{1},
	},
}

func TestSortColors(t *testing.T) {

	for _, tt := range tests {
		sortColors(tt.arg1)
		if !reflect.DeepEqual(tt.arg1, tt.want) {
			t.Errorf("got = %v, want = %v", tt.arg1, tt.want)
		}
	}
}

func TestSortColorsQuickSort(t *testing.T) {

	for _, tt := range tests {
		sortColorsQuickSort(tt.arg1)
		if !reflect.DeepEqual(tt.arg1, tt.want) {
			t.Errorf("got = %v, want = %v", tt.arg1, tt.want)
		}
	}
}

func BenchmarkSortColors(b *testing.B) {
	arg1 := []int{2, 7, 11, 15}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortColors(arg1)
	}
}

func BenchmarkSortColorsQuickSort(b *testing.B) {
	arg1 := []int{2, 7, 11, 15}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortColorsQuickSort(arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0075.Sort-Colors -bench=.
BenchmarkSortColors-8                   157253976                8.111 ns/op           0 B/op          0 allocs/op
BenchmarkSortColorsQuickSort-8          63078445                20.91 ns/op            0 B/op          0 allocs/op
*/
