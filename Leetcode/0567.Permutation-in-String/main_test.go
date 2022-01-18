package permutationinstring

import "testing"

var tests = []struct {
	arg1 string
	arg2 string
	want bool
}{
	{
		"ab",
		"eidbaooo",
		true,
	},
	{
		"ab",
		"eidboaoo",
		false,
	},
}

func TestCheckInclusion(t *testing.T) {
	for _, tt := range tests {
		if got := CheckInclusion(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestCheckInclusionSlice(t *testing.T) {
	for _, tt := range tests {
		if got := CheckInclusionSlice(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkCheckInclusion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckInclusion(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkCheckInclusionSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckInclusionSlice(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0567.Permutation-in-String -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0567.Permutation-in-String
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkCheckInclusion-8                9091321               167.9 ns/op             0 B/op          0 allocs/op
BenchmarkCheckInclusionSlice-8          26744336                53.28 ns/op            0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0567.Permutation-in-String      3.143s
*/
