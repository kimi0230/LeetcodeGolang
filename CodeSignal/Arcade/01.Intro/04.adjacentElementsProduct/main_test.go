package adjacentelementsproduct

import (
	"fmt"
	"testing"
)

func TestAdjacentElementEproduct(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{3, 6, -2, -5, 7, 3}, 21},
		{[]int{-1, -2}, 2},
		{[]int{5, 1, 2, 3, 1, 4}, 6},
		{[]int{9, 5, 10, 2, 24, -1, -48}, 50},
		{[]int{5, 6, -4, 2, 3, 2, -23}, 30},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("input: %v", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := adjacentElementEproduct(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
