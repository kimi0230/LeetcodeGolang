package longestconsecutivesequence

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{100, 4, 200, 1, 3, 2},
		4,
	},
}

func TestLongestConsecutive(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := longestConsecutive(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestConsecutive(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
