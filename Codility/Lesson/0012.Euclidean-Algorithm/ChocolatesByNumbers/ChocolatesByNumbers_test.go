package chocolatesbynumbers

import "testing"

var tests = []struct {
	arg1 int
	arg2 int
	want int
}{
	{
		10,
		4,
		5,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
