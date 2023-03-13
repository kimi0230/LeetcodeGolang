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
	{
		[]int{2},
		3,
		-1,
	},
	{
		[]int{1},
		0,
		0,
	},
}

func TestCoinChange(t *testing.T) {
	for _, tt := range tests {
		if got := CoinChange(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
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

func BenchmarkCoinChange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CoinChange(arg1, arg2)
	}
}

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
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkCoinChange-4                             273376              4452 ns/op               0 B/op          0 allocs/op
BenchmarkCoinChangeDP-4                         11071686               128.1 ns/op            96 B/op          1 allocs/op
BenchmarkCoinChangeMemoryTableRecursion-4       57663068                23.69 ns/op            0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0322.Coin-Change        4.194s
*/
