package binarynumberwithalternatingbits

import "testing"

var tests = []struct {
	arg1 int
	want bool
}{
	{
		5,
		true,
	},
	{
		7,
		false,
	},
	{
		11,
		false,
	},
	{
		10,
		true,
	},
}

func TestHasAlternatingBits(t *testing.T) {
	for _, tt := range tests {
		if got := hasAlternatingBits(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestHasAlternatingBits2(t *testing.T) {
	for _, tt := range tests {
		if got := hasAlternatingBits2(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkHasAlternatingBits(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasAlternatingBits(10)
	}
}

func BenchmarkHasAlternatingBits2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasAlternatingBits2(10)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/0693.Binary-Number-with-Alternating-Bits -bench=.
BenchmarkHasAlternatingBits-8           300204136                3.979 ns/op           0 B/op          0 allocs/op
BenchmarkHasAlternatingBits2-8          1000000000               0.3335 ns/op          0 B/op          0 allocs/op
*/
