package searcha2dmatrix

import "testing"

var tests = []struct {
	matrix [][]int
	target int
	want   bool
}{
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		3,
		true,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		13,
		false,
	},
	{
		[][]int{},
		5,
		false,
	},
}

func TestSearchMatrix(t *testing.T) {
	for _, tt := range tests {
		if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
			t.Errorf("searchMatrix(%v, %v) = %v, want %v", tt.matrix, tt.target, got, tt.want)
		}
	}
}

func BenchmarkSearchMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchMatrix(tests[0].matrix, tests[0].target)
	}
}
