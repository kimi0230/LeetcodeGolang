package binarysearch

import "testing"

var tests = []struct {
	arg1 []int
	arg2 int
	want int
}{
	{
		[]int{-1, 0, 3, 5, 9, 12},
		9,
		4,
	},
	{
		[]int{-1, 0, 3, 5, 9, 12},
		2,
		-1,
	},
	{
		[]int{-1, 0, 3, 5, 9, 12},
		13,
		-1,
	},
}

var leftboundtests = []struct {
	arg1 []int
	arg2 int
	want int
}{
	{
		[]int{1, 2, 2, 2, 3},
		2,
		1,
	},
	{
		[]int{2, 3, 5, 7},
		1,
		-1,
	},
	{
		[]int{2, 3, 5, 7},
		8,
		-1,
	},
}

var rightboundtests = []struct {
	arg1 []int
	arg2 int
	want int
}{
	{
		[]int{1, 2, 2, 4},
		2,
		2,
	},
	{
		[]int{1, 2, 2, 4},
		0,
		-1,
	},
	{
		[]int{2, 3, 5, 7},
		7,
		3,
	},
}

func TestSearch(t *testing.T) {
	for _, tt := range tests {
		if got := Search(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSearch2(t *testing.T) {
	for _, tt := range tests {
		if got := Search2(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestLeftBound(t *testing.T) {
	for _, tt := range leftboundtests {
		if got := LeftBound(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestRightBound(t *testing.T) {
	for _, tt := range rightboundtests {
		if got := RightBound(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range tests {
		if got := BinarySearch(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestBinarySearchRecur(t *testing.T) {
	for _, tt := range tests {
		if got := BinarySearchRecur(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(tests[0].arg1, tests[0].arg2)
	}
}
func BenchmarkSearch2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search2(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkBinarySearchRecur(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchRecur(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0704.Binary-Search -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0704.Binary-Search
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkSearch-4               365892498                3.483 ns/op           0 B/op          0 allocs/op
BenchmarkSearch2-4              405742570                2.734 ns/op           0 B/op          0 allocs/op
BenchmarkBinarySearch-4         81949826                13.15 ns/op            0 B/op          0 allocs/op
BenchmarkBinarySearchRecur-4        253851086                4.588 ns/op           0 B/op          0 allocs/op
*/
