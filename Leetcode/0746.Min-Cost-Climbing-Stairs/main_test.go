package mincostclimbingstairs

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{10, 15, 20},
		15,
	},
	{
		[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		6,
	},
}

func TestMinCostClimbingStairs(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MinCostClimbingStairs(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestMinCostClimbingStairsOptimize(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MinCostClimbingStairsOptimize(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkMinCostClimbingStairs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinCostClimbingStairs(tests[0].arg1)
	}
}

func BenchmarkMinCostClimbingStairsOptimize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinCostClimbingStairsOptimize(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0746.Min-Cost-Climbing-Stairs -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0746.Min-Cost-Climbing-Stairs
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkMinCostClimbingStairs-4                36693742                30.17 ns/op           24 B/op          1 allocs/op
BenchmarkMinCostClimbingStairsOptimize-4        405489464                3.091 ns/op           0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0746.Min-Cost-Climbing-Stairs   2.713s
*/
