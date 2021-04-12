package minimumsizesubarraysum

import (
	"testing"
)

var tests = []struct {
	arg1 int
	arg2 []int
	want int
}{
	{
		7,
		[]int{2, 3, 1, 2, 4, 3},
		2,
	},
	{
		4,
		[]int{1, 4, 4},
		1,
	},
	{
		11,
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		0,
	},
}

func TestMinSubArrayLenBurst(t *testing.T) {
	for _, tt := range tests {
		if got := MinSubArrayLenBurst(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func TestMinSubArrayLenSlidingWindow(t *testing.T) {
	for _, tt := range tests {
		if got := MinSubArrayLenSlidingWindow(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func TestMinSubArrayLenBinarysearch(t *testing.T) {
	for _, tt := range tests {
		if got := MinSubArrayLenBinarysearch(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}

func BenchmarkMinSubArrayLenBurst(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinSubArrayLenBurst(tests[0].arg1, tests[0].arg2)
	}
}
func BenchmarkMinSubArrayLenSlidingWindow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinSubArrayLenSlidingWindow(tests[0].arg1, tests[0].arg2)
	}
}
func BenchmarkMinSubArrayLenBinarysearch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinSubArrayLenBinarysearch(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/0209.Minimum-Size-Subarray-Sum -bench=.
BenchmarkMinSubArrayLenBurst-4                  27345652                41.67 ns/op            0 B/op          0 allocs/op
BenchmarkMinSubArrayLenSlidingWindow-4          87522610                13.67 ns/op            0 B/op          0 allocs/op
BenchmarkMinSubArrayLenBinarysearch-4            8496684               136.7 ns/op            64 B/op          1 allocs/op
*/
