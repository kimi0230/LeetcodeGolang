package twosumiiinputarrayissorted

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	arg2 int
	want []int
}{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{1, 2},
	},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range tests {
		if got := TwoSum(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	for _, tt := range tests {
		if got := TwoSum2(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0167.Two-Sum-II-Input-Array-Is-Sorted -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0167.Two-Sum-II-Input-Array-Is-Sorted
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkTwoSum-4       35161128                39.74 ns/op           16 B/op          1 allocs/op
BenchmarkTwoSum2-4      19309680                70.21 ns/op           16 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0167.Two-Sum-II-Input-Array-Is-Sorted   2.866s
*/
