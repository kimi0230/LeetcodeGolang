package besttimetobuyandsellstock

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{7, 1, 5, 3, 6, 4},
		5,
	},
	{
		[]int{7, 6, 4, 3, 1},
		0,
	},
}

func TestMaxProfit(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MaxProfit(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestMaxProfitDP(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MaxProfitDP(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestMaxProfitDP2(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MaxProfitDP2(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfit(tests[0].arg1)
	}
}

func BenchmarkMaxProfitDP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfitDP(tests[0].arg1)
	}
}

func BenchmarkMaxProfitDP2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfitDP2(tests[0].arg1)
	}
}

/*
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0121.Best-Time-to-Buy-and-Sell-Stock
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMaxProfit-8            208734571                8.156 ns/op           0 B/op          0 allocs/op
BenchmarkMaxProfitDP-8          240880467                5.720 ns/op           0 B/op          0 allocs/op
BenchmarkMaxProfitDP2-8          4095866               282.6 ns/op           240 B/op          7 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0121.Best-Time-to-Buy-and-Sell-Stock    6.732s
*/
