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

*/
