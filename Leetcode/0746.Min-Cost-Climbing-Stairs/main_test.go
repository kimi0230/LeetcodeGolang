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

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
