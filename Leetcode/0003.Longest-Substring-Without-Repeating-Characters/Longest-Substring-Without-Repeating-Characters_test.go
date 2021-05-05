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

func TestLengthOfLongestSubstringMap(t *testing.T) {
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
		if got := LengthOfLongestSubstringMap(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLengthOfLongestSubstringBit(t *testing.T) {
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
pkg: LeetcodeGolang/0003.Longest-Substring-Without-Repeating-Characters
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkLengthOfLongestSubstring-4             51885124                20.68 ns/op            0 B/op          0 allocs/op
BenchmarkLengthOfLongestSubstringMap-4           2677759               424.5 ns/op             0 B/op          0 allocs/op
BenchmarkLengthOfLongestSubstringBit-4          54287119                18.82 ns/op            0 B/op          0 allocs/op
*/
