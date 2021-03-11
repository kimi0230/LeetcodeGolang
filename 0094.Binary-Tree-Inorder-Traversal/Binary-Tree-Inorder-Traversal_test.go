package binarytreeinordertraversal

import (
	"LeetcodeGolang/Utility/structures"
	"reflect"
	"testing"
)

func TestInorderTraversalStack(t *testing.T) {

	tests := []struct {
		arg1 []int
		want []int
	}{
		{
			arg1: []int{1, structures.NULL, 2, 3},
			want: []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		root := structures.Ints2TreeNode(tt.arg1)
		if got := InorderTraversalStack(root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", tt.arg1, tt.want)
		}
	}
}

func TestInorderTraversalRecursion(t *testing.T) {

	tests := []struct {
		arg1 []int
		want []int
	}{
		{
			arg1: []int{1, structures.NULL, 2, 3},
			want: []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		root := structures.Ints2TreeNode(tt.arg1)
		if got := InorderTraversalRecursion(root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", tt.arg1, tt.want)
		}
	}
}

func BenchmarkInorderTraversalStack(b *testing.B) {
	arg1 := []int{1, structures.NULL, 2, 3}
	root := structures.Ints2TreeNode(arg1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InorderTraversalStack(root)
	}
}

func BenchmarkInorderTraversalRecursion(b *testing.B) {
	arg1 := []int{1, structures.NULL, 2, 3}
	root := structures.Ints2TreeNode(arg1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InorderTraversalRecursion(root)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/0094.Binary-Tree-Inorder-Traversal -bench=.
BenchmarkInorderTraversalStack-8         4469162               267.9 ns/op           568 B/op          4 allocs/op
BenchmarkInorderTraversalRecursion-8    11604469               103.4 ns/op            56 B/op          3 allocs/op
*/
