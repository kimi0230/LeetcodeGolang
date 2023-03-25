package invertbinarytree

import (
	"LeetcodeGolang/Utility/structures"
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	want []int
}{
	{
		[]int{4, 2, 7, 1, 3, 6, 9},
		[]int{4, 7, 2, 9, 6, 3, 1},
	},
}

func TestInvertTree(t *testing.T) {
	for _, tt := range tests {
		tree := structures.IntsToTree(tt.arg1)
		got := InvertTree(tree)
		treeWant := structures.IntsToTree(tt.want)
		if !reflect.DeepEqual(got, treeWant) {
			t.Errorf("got = %v, want = %v", got, treeWant)
		}
	}
}

func BenchmarkInvertTree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree := structures.IntsToTree(tests[0].arg1)
		InvertTree(tree)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0226.Invert-Binary-Tree -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0226.Invert-Binary-Tree
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkInvertTree-4            2602960               532.4 ns/op           168 B/op          7 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0226.Invert-Binary-Tree 1.869s
*/
