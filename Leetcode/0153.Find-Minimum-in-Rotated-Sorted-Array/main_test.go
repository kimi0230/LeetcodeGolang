package findminimuminrotatedsortedarray

import "testing"

var tests = []struct {
	nums []int
	want int
}{
	{[]int{3, 4, 5, 1, 2}, 1},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	{[]int{11, 13, 15, 17}, 11},
	{[]int{2, 1}, 1},
	{[]int{1}, 1},
}

func TestFindMin(t *testing.T) {
	for _, tt := range tests {
		if got := findMin(tt.nums); got != tt.want {
			t.Errorf("findMin(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func BenchmarkFindMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMin(tests[0].nums)
	}
}
