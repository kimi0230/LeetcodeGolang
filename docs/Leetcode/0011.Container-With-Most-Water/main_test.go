package containerwithmostwater

import "testing"

var tests = []struct {
	height []int
	want   int
}{
	{
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		49,
	},
	{
		[]int{4, 3, 2, 1, 4},
		16,
	},
	{
		[]int{1, 2, 1},
		2,
	},
	{
		[]int{1, 1},
		1,
	},
}

func TestMaxArea(t *testing.T) {
	for _, tt := range tests {
		if got := maxArea(tt.height); got != tt.want {
			t.Errorf("maxArea(%v) = %v, want %v", tt.height, got, tt.want)
		}
	}
}

func BenchmarkMaxArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxArea(tests[0].height)
	}
}
