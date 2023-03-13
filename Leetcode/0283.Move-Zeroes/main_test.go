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

func BenchmarkMoveZeroes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MoveZeroes(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
