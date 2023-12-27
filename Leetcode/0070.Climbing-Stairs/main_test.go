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
}

func TestClimbStairs(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := ClimbStairs(tt.arg1); got != tt.want {
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
		ClimbStairs(tests[0].arg1)
	}
}

func BenchmarkClimbStairsRecursive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ClimbStairsRecursive(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0070.Climbing-Stairs -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0070.Climbing-Stairs
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkClimbStairs-8                  40758399                30.75 ns/op           24 B/op          1 allocs/op
BenchmarkClimbStairsRecursive-8         42196260                27.60 ns/op           24 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0070.Climbing-Stairs    2.489s
*/
