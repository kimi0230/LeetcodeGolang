package longestSubstringwithoutrepeatingcharacters

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 string
	want int
}{
	{
		arg1: "abcabcbb",
		want: 3,
	},
	{
		arg1: "bbbbbb",
		want: 1,
	},
	{
		arg1: "pwwkew",
		want: 3,
	},
	{
		arg1: "",
		want: 0,
	},
	{
		arg1: "aab",
		want: 2,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tt := range tests {
		if got := LengthOfLongestSubstring(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLengthOfLongestSubstringMap(t *testing.T) {
	for _, tt := range tests {
		if got := LengthOfLongestSubstringMap(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLengthOfLongestSubstringBit(t *testing.T) {
	for _, tt := range tests {
		if got := LengthOfLongestSubstringBit(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	arg1 := "abcabcbb"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstring(arg1)
	}
}

func BenchmarkLengthOfLongestSubstringMap(b *testing.B) {
	arg1 := "abcabcbb"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstringMap(arg1)
	}
}

func BenchmarkLengthOfLongestSubstringBit(b *testing.B) {
	arg1 := "abcabcbb"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstringBit(arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0003.Longest-Substring-Without-Repeating-Characters -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0003.Longest-Substring-Without-Repeating-Characters
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkLengthOfLongestSubstring-8             67657353                19.30 ns/op            0 B/op          0 allocs/op
BenchmarkLengthOfLongestSubstringMap-8           2463084               483.5 ns/op             0 B/op          0 allocs/op
BenchmarkLengthOfLongestSubstringBit-8          53792877                21.70 ns/op            0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0003.Longest-Substring-Without-Repeating-Characters     5.125s
*/
