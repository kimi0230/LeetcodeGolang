package sametree

import (
	"LeetcodeGolang/structures"
	"testing"
)

var tests = []struct {
	arg1 []int
	arg2 []int
	want bool
}{
	{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		true,
	},
	{
		[]int{1, 2},
		[]int{1, structures.NULL, 2},
		false,
	},
	{
		[]int{1, 2, 1},
		[]int{1, 1, 2},
		false,
	},
}

func TestIsSameTree(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		p := structures.Ints2TreeNode(tt.arg1)
		q := structures.Ints2TreeNode(tt.arg2)

		if got := IsSameTree(p, q); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkIsSameTree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := structures.Ints2TreeNode(tests[0].arg1)
		q := structures.Ints2TreeNode(tests[0].arg2)
		IsSameTree(p, q)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
