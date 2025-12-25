package maximumdepthofbinarytree

import (
	"LeetcodeGolang/structures"
	"testing"
)

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{3, 9, 20, 0, 0, 15, 7},
		3,
	},
}

func TestMaxDepth(t *testing.T) {
	for _, tt := range tests {
		root := structures.Ints2TreeNode(tt.arg1)
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MaxDepth(root); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestMaxDepthKimi(t *testing.T) {
	for _, tt := range tests {
		root := structures.Ints2TreeNode(tt.arg1)
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MaxDepthKimi(root); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := structures.Ints2TreeNode(tests[0].arg1)
		MaxDepth(root)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
