package maxareaofisland

import "testing"

var tests = []struct {
	arg1 [][]int
	want int
}{
	{
		[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
		6,
	},
	{
		[][]int{{0, 0, 0, 0, 0, 0, 0, 0}},
		0,
	},
}

func TestMaxAreaOfIsland(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MaxAreaOfIsland(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkMaxAreaOfIsland(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxAreaOfIsland(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0695.Max-Area-of-Island -bench=.

*/
