package floodfill

import (
	"reflect"
	"testing"
)

var tests = []struct {
	image         [][]int
	sr, sc, color int
	expected      [][]int
}{
	{
		[][]int{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
		},
		1, 1, 2,
		[][]int{
			{2, 2, 2},
			{2, 2, 0},
			{2, 0, 1},
		},
	},
	{
		[][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
		1, 1, 1,
		[][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	},
	{
		[][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
		0, 0, 0,
		[][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	},
	// Add more test cases here if needed
}

func TestFloodFillDFS(t *testing.T) {
	for _, tt := range tests {
		result := FloodFill(tt.image, tt.sr, tt.sc, tt.color)
		// Check if the result matches the expected output
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("FloodFillDFS(%v, %d, %d, %d) returned %+v, expected %+v",
				tt.image, tt.sr, tt.sc, tt.color, result, tt.expected)
		}
	}
}

func TestFloodFillBFS(t *testing.T) {
	for _, tt := range tests {
		result := FloodFillBFS(tt.image, tt.sr, tt.sc, tt.color)
		// Check if the result matches the expected output
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("FloodFillBFS(%v, %d, %d, %d) returned %+v, expected %+v",
				tt.image, tt.sr, tt.sc, tt.color, result, tt.expected)
		}
	}
}

func BenchmarkFloodFillDFS(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FloodFill(tests[0].image, tests[0].sr, tests[0].sc, tests[0].color)
	}
}

func BenchmarkFloodFillBFS(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FloodFillBFS(tests[0].image, tests[0].sr, tests[0].sc, tests[0].color)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0733.Flood-Fill -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0733.Flood-Fill
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkFloodFillDFS-8         384756555                4.486 ns/op           0 B/op          0 allocs/op
BenchmarkFloodFillBFS-8         309088303                3.642 ns/op           0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0733.Flood-Fill 3.572s
*/
