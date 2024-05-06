package openthelock

import (
	"testing"
)

var tests = []struct {
	arg1 []string
	arg2 string
	want int
}{
	{
		[]string{"0201", "0101", "0102", "1212", "2002"},
		"0202",
		6,
	},
	{
		[]string{"8888"},
		"0009",
		1,
	},
	{
		[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
		"8888",
		-1,
	},
}

func TestOpenLock(t *testing.T) {
	for _, tt := range tests {
		if got := OpenLock(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func TestOpenLockBiDirection(t *testing.T) {
	for _, tt := range tests {
		if got := OpenLockBiDirection(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func TestOpenLockBiDirectionＯptimization(t *testing.T) {
	for _, tt := range tests {
		if got := OpenLockBiDirectionOptimization(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func BenchmarkOpenLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OpenLock(tests[i%3].arg1, tests[i%3].arg2)
	}
}

func BenchmarkOpenLockBiDirection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OpenLockBiDirection(tests[i%3].arg1, tests[i%3].arg2)
	}
}

func BenchmarkOpenLockBiDirectionOptimization(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OpenLockBiDirectionOptimization(tests[i%3].arg1, tests[i%3].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0752.Open-the-Lock -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0752.Open-the-Lock
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkOpenLock-8                                12866             92523 ns/op            5543 B/op         12 allocs/op
BenchmarkOpenLockBiDirection-8                     79819             14357 ns/op            1722 B/op         28 allocs/op
BenchmarkOpenLockBiDirectionOptimization-8         85179             14486 ns/op            1721 B/op         28 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0752.Open-the-Lock      4.820s
*/
