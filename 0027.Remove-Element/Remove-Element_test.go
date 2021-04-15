package removeelement

import (
	"testing"
)

var tests = []struct {
	arg1 []int
	arg2 int
	want int
}{
	{
		[]int{1, 0, 1},
		1,
		1,
	},
	{
		[]int{0, 1, 0, 3, 0, 12},
		0,
		3,
	},
	{
		[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 12},
		0,
		4,
	},
	{
		[]int{0, 0, 0, 0, 0},
		0,
		0,
	},

	{
		[]int{1},
		1,
		0,
	},
	{
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
		2,
		5,
	},
}

func TestRemoveElementDoublePoint(t *testing.T) {
	for _, tt := range tests {
		if got := RemoveElementDoublePoint(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func TestRemoveElement(t *testing.T) {
	for _, tt := range tests {
		if got := RemoveElement(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElement(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkRemoveElementDoublePoint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElementDoublePoint(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/0027.Remove-Element -bench=.
BenchmarkRemoveElement-8                100000000               10.66 ns/op            0 B/op          0 allocs/op
BenchmarkRemoveElementDoublePoint-8     285200671                5.456 ns/op           0 B/op          0 allocs/op
*/
