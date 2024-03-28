package climbingstairs

import "testing"

var tests = []struct {
	arg1 int
	want int
}{
	{
		2,
		2,
	},
	{
		3,
		3,
	},
	{
		39,
		63245986,
	},
}

func TestClimbStairs(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := ClimbStairs(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestClimbStairsCache(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := ClimbStairsCache(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestClimbStairsRecursive(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := ClimbStairsRecursive(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ClimbStairs(tests[2].arg1)
	}
}

func BenchmarkClimbStairsCache(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ClimbStairsCache(tests[2].arg1)
	}
}

func BenchmarkClimbStairsRecursive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ClimbStairsRecursive(tests[2].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0070.Climbing-Stairs -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0070.Climbing-Stairs
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkClimbStairs-8                  10386211               112.1 ns/op           320 B/op          1 allocs/op
BenchmarkClimbStairsCache-8             10184984               118.8 ns/op           320 B/op          1 allocs/op
BenchmarkClimbStairsRecursive-8                4         281980486 ns/op             320 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0070.Climbing-Stairs    5.591s
*/
