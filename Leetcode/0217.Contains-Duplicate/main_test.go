package containsduplicate

import "testing"

var tests = []struct {
	input    []int
	expected bool
}{
	{[]int{1, 2, 3, 1}, true},                   // Duplicate element present
	{[]int{1, 2, 3, 4}, false},                  // No duplicate element
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true}, // Duplicate element present
}

func TestContainsDuplicate(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.expected) {
		if got := ContainsDuplicate(tt.input); got != tt.expected {
			t.Errorf("got = %v, expected = %v", got, tt.expected)
		}
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicate(tests[0].input)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
