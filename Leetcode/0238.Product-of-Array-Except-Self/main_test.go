package productofarrayexceptself

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	want []int
}{
	{
		[]int{1, 2, 3, 4},
		[]int{24, 12, 8, 6},
	},
}

func TestProductExceptSelf(t *testing.T) {
	for _, tt := range tests {
		if got := productExceptSelf(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			// if got := productExceptSelf(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		productExceptSelf(tests[0].arg1)
	}
}

func TestProductExceptSelf2(t *testing.T) {
	for _, tt := range tests {
		if got := productExceptSelf2(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			// if got := productExceptSelf2(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkProductExceptSelf2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		productExceptSelf2(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0238.Product-of-Array-Except-Self -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0238.Product-of-Array-Except-Self
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkProductExceptSelf-8            12772174               158.4 ns/op            96 B/op          3 allocs/op
BenchmarkProductExceptSelf2-8           32292304                63.74 ns/op           32 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0238.Product-of-Array-Except-Self       4.228s
*/
