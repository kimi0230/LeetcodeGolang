package validanagram

import "testing"

var tests = []struct {
	s        string
	t        string
	expected bool
}{
	{"anagram", "nagaram", true}, // 是有效的字母異位詞
	{"rat", "car", false},        // 不是有效的字母異位詞
	{"listen", "silent", true},   // 是有效的字母異位詞
	{"hello", "world", false},    // 不是有效的字母異位詞
}

func TestIsAnagram(t *testing.T) {

	for _, test := range tests {
		result := IsAnagram(test.s, test.t)
		if result != test.expected {
			t.Errorf("s: %s, t: %s, Expected: %v, Got: %v", test.s, test.t, test.expected, result)
		}
	}
}

func TestIsAnagram2(t *testing.T) {

	for _, test := range tests {
		result := IsAnagram2(test.s, test.t)
		if result != test.expected {
			t.Errorf("s: %s, t: %s, Expected: %v, Got: %v", test.s, test.t, test.expected, result)
		}
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	s := "anagram"
	t := "nagaram"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsAnagram(s, t)
	}
}

func BenchmarkIsAnagram2(b *testing.B) {
	s := "anagram"
	t := "nagaram"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsAnagram2(s, t)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0242.Valid-Anagram -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0242.Valid-Anagram
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkIsAnagram-8             6703497               276.3 ns/op             0 B/op          0 allocs/op
BenchmarkIsAnagram2-8            3660909               318.9 ns/op            48 B/op          2 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0242.Valid-Anagram      4.498s
*/
