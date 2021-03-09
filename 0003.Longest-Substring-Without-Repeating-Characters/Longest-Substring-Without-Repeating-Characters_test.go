package longestSubstringwithoutrepeatingcharacters

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
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
	}

	for _, tt := range tests {
		if got := LengthOfLongestSubstring(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLengthOfLongestSubstringPin(t *testing.T) {
	tests := []struct {
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
	}

	for _, tt := range tests {
		if got := LengthOfLongestSubstringPin(tt.arg1); !reflect.DeepEqual(got, tt.want) {
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

func BenchmarkLengthOfLongestSubstringPin(b *testing.B) {
	arg1 := "abcabcbb"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstringPin(arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/0003.Longest-Substring-Without-Repeating-Characters -bench=.
BenchmarkLengthOfLongestSubstring-8             62972542                23.09 ns/op            0 B/op          0 allocs/op
BenchmarkLengthOfLongestSubstringPin-8           3590331               354.7 ns/op             0 B/op          0 allocs/op
*/
