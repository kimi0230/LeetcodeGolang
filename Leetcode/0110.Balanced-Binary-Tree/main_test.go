package balancedbinarytree

import (
	"LeetcodeGolang/structures"
	"testing"
)

var tests = []struct {
	arg1 []int
	want bool
}{
	{
		[]int{1, 2, 2, 3, structures.NULL, structures.NULL, 3, 4, structures.NULL, structures.NULL, 4},
		false,
	},
	{
		[]int{1, 2, 2, 3, structures.NULL, structures.NULL, 3, 4, structures.NULL, structures.NULL, 4},
		false,
	},
	{
		[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7},
		true,
	},
}

func TestIsBalanced(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		root := structures.Ints2TreeNode(tt.arg1)

		if got := IsBalanced(root); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := structures.Ints2TreeNode(tests[0].arg1)
		IsBalanced(root)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
