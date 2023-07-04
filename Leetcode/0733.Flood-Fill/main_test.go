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

func TestFloodFill(t *testing.T) {
	for _, tt := range tests {
		result := FloodFill(tt.image, tt.sr, tt.sc, tt.color)
		// Check if the result matches the expected output
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("FloodFill(%v, %d, %d, %d) returned %+v, expected %+v",
				tt.image, tt.sr, tt.sc, tt.color, result, tt.expected)
		}
	}
}

func BenchmarkFloodFill(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FloodFill(tests[0].image, tests[0].sr, tests[0].sc, tests[0].color)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
