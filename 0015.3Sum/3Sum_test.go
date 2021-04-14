package threesum

import (
	"reflect"
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

/*
go test -benchmem -run=none LeetcodeGolang/0015.3Sum -bench=.
BenchmarkThreeSumBurst-8                 8780395               160.7 ns/op            72 B/op          3 allocs/op
BenchmarkThreeSumDoublePoint-8          10466467               127.8 ns/op            72 B/op          3 allocs/op
*/
