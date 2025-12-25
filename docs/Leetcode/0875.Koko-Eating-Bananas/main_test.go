package kokoeatingbananas

import "testing"

var tests = []struct {
	piles []int
	h     int
	want  int
}{
	{
		[]int{3, 6, 7, 11},
		8,
		4,
	},
	{
		[]int{30, 11, 23, 4, 20},
		5,
		30,
	},
	{
		[]int{30, 11, 23, 4, 20},
		6,
		23,
	},
}

func TestMinEatingSpeed(t *testing.T) {
	for _, tt := range tests {
		if got := minEatingSpeed(tt.piles, tt.h); got != tt.want {
			t.Errorf("minEatingSpeed(%v, %v) = %v, want %v", tt.piles, tt.h, got, tt.want)
		}
	}
}

func BenchmarkMinEatingSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minEatingSpeed(tests[0].piles, tests[0].h)
	}
}
