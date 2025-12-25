package movezeroes

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	want []int
}{
	{
		[]int{0, 1, 0, 3, 12},
		[]int{1, 3, 12, 0, 0},
	},
	{
		[]int{0},
		[]int{0},
	},
}

func TestMoveZeroes(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if MoveZeroes(tt.arg1); !reflect.DeepEqual(tt.arg1, tt.want) {
			t.Errorf("got = %v, want = %v", tt.arg1, tt.want)
		}
	}
}

func TestMoveZeroes2point(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if MoveZeroes2point(tt.arg1); !reflect.DeepEqual(tt.arg1, tt.want) {
			t.Errorf("got = %v, want = %v", tt.arg1, tt.want)
		}
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MoveZeroes(tests[0].arg1)
	}
}

func BenchmarkMoveZeroes2point(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MoveZeroes2point(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0283.Move-Zeroes -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0283.Move-Zeroes
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkMoveZeroes-4           145544966                9.164 ns/op           0 B/op          0 allocs/op
BenchmarkMoveZeroes2point-4     183269922                6.149 ns/op           0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0283.Move-Zeroes        3.975s
*/
