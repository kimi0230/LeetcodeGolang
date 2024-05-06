// TODO : 0002.Add-Two-Numbers Unit Test
package 

import "testing"

var tests = []struct {
	arg1 string
	want int
}{
	{
		"bbbab",
		4,
	},
}

func TestLongestPalindromeSubseq(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := LongestPalindromeSubseq(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkLongestPalindromeSubseq(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestPalindromeSubseq(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
