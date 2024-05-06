package findminimumdifferencebetweenanytwoelements

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{1, 5, 3, 19, 18, 25},
		1,
	},
	{
		[]int{30, 5, 20, 9},
		4,
	},
	{
		[]int{1, 19, -4, 31, 38, 25, 100},
		5,
	},
	{
		[]int{1},
		0,
	},
}

func TestFindMinDiff(t *testing.T) {
	for _, tt := range tests {
		if got := FindMinDiff(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkFindMinDiff(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMinDiff(tests[0].arg1)
	}
}

func BenchmarkFindMinDiff2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMinDiff(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/GeeksforGeeks/SortingAlgorithms/0031.Find-Minimum-Difference-Between-Any-Two-Elements -bench=.
*/
