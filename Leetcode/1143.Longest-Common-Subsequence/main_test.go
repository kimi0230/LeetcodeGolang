package longestcommonsubsequence

import (
	"testing"
)

var tests = []struct {
	arg1 string
	arg2 string
	want int
}{
	{
		"abcde",
		"ace",
		3,
	},
	{
		"abc",
		"abc",
		3,
	},
	{
		"abc",
		"def",
		0,
	},
	{
		"pmjghexybyrgzczy",
		"hafcdqbgncrcbihkd",
		4,
	},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, tt := range tests {
		if got := LongestCommonSubsequence(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLongestCommonSubsequenceDP(t *testing.T) {
	for _, tt := range tests {
		if got := LongestCommonSubsequenceDP(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestCommonSubsequence(tests[i%4].arg1, tests[i%4].arg2)
	}
}

func BenchmarkLongestCommonSubsequenceDP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestCommonSubsequenceDP(tests[i%4].arg1, tests[i%4].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/1143.Longest-Common-Subsequence -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/1143.Longest-Common-Subsequence
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkLongestCommonSubsequence-8                  100         737158262 ns/op               0 B/op          0 allocs/op
BenchmarkLongestCommonSubsequenceDP-8            2355297               491.3 ns/op           912 B/op          8 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/1143.Longest-Common-Subsequence 75.400s
*/
