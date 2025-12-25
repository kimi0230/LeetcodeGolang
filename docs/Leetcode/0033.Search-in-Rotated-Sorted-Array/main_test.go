package searchinrotatedsortedarray

import (
	"testing"
)

var tests = []struct {
	nums   []int
	target int
	want   int
}{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	{[]int{1}, 0, -1},
}

func TestSearch(t *testing.T) {
	for _, tt := range tests {
		if got := search(tt.nums, tt.target); got != tt.want {
			t.Errorf("search(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			search(tt.nums, tt.target)
		}
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
