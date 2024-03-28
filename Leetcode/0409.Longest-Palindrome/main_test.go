package longestpalindrome

import "testing"

var tests = []struct {
	arg1 string
	want int
}{
	{
		"abccccdd",
		7,
	},
	{
		"a",
		1,
	},
}

func TestLongestPalindrome(t *testing.T) {
	for _, tt := range tests {
		if got := LongestPalindrome(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestPalindrome(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
