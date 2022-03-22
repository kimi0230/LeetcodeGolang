package editdistance

import "testing"

var tests = []struct {
	arg1 string
	arg2 string
	want int
}{
	{
		"horse",
		"ros",
		3,
	},
	{
		"intention",
		"execution",
		5,
	},
}

func TestMinDistance(t *testing.T) {
	for _, tt := range tests {
		if got := MinDistance(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestMinDistanceMemo(t *testing.T) {
	for _, tt := range tests {
		if got := MinDistanceMemo(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkMinDistance(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinDistance(tests[i%2].arg1, tests[i%2].arg2)
	}
}

func BenchmarkMinDistanceMemo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinDistanceMemo(tests[i%2].arg1, tests[i%2].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0072.Edit-Distance -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0072.Edit-Distance
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMinDistance-8            415807              3888 ns/op               0 B/op          0 allocs/op
BenchmarkMinDistanceMemo-8         64245             16935 ns/op            2212 B/op         69 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0072.Edit-Distance      2.940s
*/
