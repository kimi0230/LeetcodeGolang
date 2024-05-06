package findallanagramsinastring

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 string
	arg2 string
	want []int
}{
	{
		"cbaebabacd",
		"abc",
		[]int{0, 6},
	},
	{
		"abab",
		"ab",
		[]int{0, 1, 2},
	},
}

func TestFindAnagrams(t *testing.T) {
	for _, tt := range tests {
		if got := FindAnagrams(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestFindAnagramsSlice(t *testing.T) {
	for _, tt := range tests {
		if got := FindAnagramsSlice(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkFindAnagrams(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindAnagrams(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkFindAnagramsSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindAnagramsSlice(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0438.Find-All-Anagrams-in-a-String -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0438.Find-All-Anagrams-in-a-String
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkFindAnagrams-8          1572444               690.4 ns/op            24 B/op          2 allocs/op
BenchmarkFindAnagramsSlice-8    11687409               136.9 ns/op            24 B/op          2 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0438.Find-All-Anagrams-in-a-String      4.329s
*/
