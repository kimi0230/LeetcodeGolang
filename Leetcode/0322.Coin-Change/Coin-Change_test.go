package coinchange

import (
	"testing"
)

var tests = []struct {
	arg1 []int
	arg2 int
	want int
}{
	{
		[]int{1, 2, 5},
		11,
		3,
	},
}

func TestCoinChangeDP(t *testing.T) {
	for _, tt := range tests {
		if got := CoinChangeDP(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func TestCoinChangeMemoryTableRecursion(t *testing.T) {
	for _, tt := range tests {
		if got := CoinChangeMemoryTableRecursion(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

var arg1 = []int{1, 2, 5}
var arg2 = 11

func BenchmarkCoinChangeDP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CoinChangeDP(arg1, arg2)
	}
}

func BenchmarkCoinChangeMemoryTableRecursion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CoinChangeMemoryTableRecursion(arg1, arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0322.Coin-Change -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0322.Coin-Change
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkCoinChangeDP-8                         13531378                91.60 ns/op           96 B/op          1 allocs/op
BenchmarkCoinChangeMemoryTableRecursion-8       50481345                21.66 ns/op            0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0322.Coin-Change        2.459s
*/
