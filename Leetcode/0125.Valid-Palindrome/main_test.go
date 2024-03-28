package validpalindrome

import "testing"

var tests = []struct {
	arg1 string
	want bool
}{
	{
		"A man, a plan, a canal: Panama",
		true,
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := IsPalindrome(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(tests[0].arg1)
	}
}

func BenchmarkIsPalindromeStrBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeStrBuilder(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0125.Valid-Palindrome -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0125.Valid-Palindrome
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkIsPalindrome-8                  3840048               552.2 ns/op            32 B/op          1 allocs/op
BenchmarkIsPalindromeStrBuilder-8        3164848               439.0 ns/op            88 B/op          4 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0125.Valid-Palindrome   5.242s
*/
