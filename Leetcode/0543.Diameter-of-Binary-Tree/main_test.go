package diameterofbinarytree

import (
	"LeetcodeGolang/structures"
	"testing"
)

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{1, 2, 3, 4, 5},
		3,
	},
	{
		[]int{1, 2},
		1,
	},
}

func TestDiameterOfBinaryTree(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		root := structures.Ints2TreeNode(tt.arg1)
		if got := DiameterOfBinaryTree(root); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkDiameterOfBinaryTree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := structures.Ints2TreeNode(tests[0].arg1)
		DiameterOfBinaryTree(root)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
