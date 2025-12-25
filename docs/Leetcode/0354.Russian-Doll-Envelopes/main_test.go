package russiandollenvelopes

import "testing"

var tests = []struct {
	arg1 [][]int
	want int
}{
	{
		[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
		3,
	},
	{
		[][]int{{1, 1}, {1, 1}, {1, 1}},
		1,
	},
}

func TestMaxEnvelopes(t *testing.T) {
	for _, tt := range tests {
		if got := MaxEnvelopes(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestMaxEnvelopes2(t *testing.T) {
	for _, tt := range tests {
		if got := MaxEnvelopes2(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkMaxEnvelopes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxEnvelopes(tests[0].arg1)
	}
}

func BenchmarkMaxEnvelopes2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxEnvelopes2(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMaxEnvelopes-8          9067520               167.0 ns/op            88 B/op          3 allocs/op
BenchmarkMaxEnvelopes2-8         6726646               214.9 ns/op            80 B/op          4 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes     6.237s
*/
