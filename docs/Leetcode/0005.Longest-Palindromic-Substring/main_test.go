package longestpalindromicsubstring

import "testing"

var tests = []struct {
	arg1 string
	want string
}{
	{
		"babad",
		"bab",
	},
	{
		"cbbd",
		"bb",
	},
}

func TestLongestPalindromeSubstring(t *testing.T) {
	for _, tt := range tests {
		if got := longestPalindrome(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkLongestPalindromeSubstring(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
