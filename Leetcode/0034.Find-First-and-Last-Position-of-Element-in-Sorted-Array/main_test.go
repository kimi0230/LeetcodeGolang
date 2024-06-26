package findfirstandlastpositionofelementinsortedarray

import "testing"

var tests = []struct {
	nums   []int
	target int
	want   []int
}{
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{}, 0, []int{-1, -1}},
	{[]int{1}, 1, []int{0, 0}},
}

func TestSearchRange(t *testing.T) {
	for _, tt := range tests {
		got := searchRange(tt.nums, tt.target)
		if got[0] != tt.want[0] || got[1] != tt.want[1] {
			t.Errorf("searchRange(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func BenchmarkSearchRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			searchRange(tt.nums, tt.target)
		}
	}
}
