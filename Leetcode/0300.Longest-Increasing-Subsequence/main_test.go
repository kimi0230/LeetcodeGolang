package longestincreasingsubsequence

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{10, 9, 2, 5, 3, 7, 101, 18},
		4,
	},
	{
		[]int{0, 1, 0, 3, 2, 3},
		4,
	},
	{
		[]int{7, 7, 7, 7, 7, 7, 7},
		1,
	},
	{
		[]int{6, 3, 5, 10, 11, 2, 9, 14, 13, 7, 4, 8, 12},
		5,
	},
}

func TestLengthOfLIS(t *testing.T) {
	for _, tt := range tests {
		if got := LengthOfLIS(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLengthOfLIS2(t *testing.T) {
	for _, tt := range tests {
		if got := LengthOfLIS2(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLengthOfLISPatience(t *testing.T) {
	for _, tt := range tests {
		if got := LengthOfLISPatience(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLIS(tests[0].arg1)
	}
}
func BenchmarkLengthOfLIS2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLIS2(tests[0].arg1)
	}
}
func BenchmarkLengthOfLISPatience(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLISPatience(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/300.Longest-Increasing-Subsequence -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/300.Longest-Increasing-Subsequence
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkLengthOfLIS-8                  18300901                64.77 ns/op           64 B/op          1 allocs/op
BenchmarkLengthOfLIS2-8                 17410562                68.98 ns/op           80 B/op          1 allocs/op
BenchmarkLengthOfLISPatience-8          20768851                58.47 ns/op           64 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/300.Longest-Increasing-Subsequence      3.812s
*/
