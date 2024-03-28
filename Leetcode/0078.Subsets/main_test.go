package subsets

import (
	"reflect"
	"sort"
	"testing"
)

var tests = []struct {
	arg1 []int
	want [][]int
}{
	{
		[]int{1, 2, 3},
		[][]int{
			{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3},
		},
	},
	{
		[]int{0},
		[][]int{
			{}, {0},
		},
	},
}

func TestSubsets(t *testing.T) {
	for _, tt := range tests {
		got := Subsets(tt.arg1)
		sortOutput(&got)
		if !reflect.DeepEqual(got, tt.want) {
			// 	// if got := Subsets(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

//go:cover
func TestSubsetsDFS(t *testing.T) {
	for _, tt := range tests {
		got := SubsetsDFS(tt.arg1)
		sortOutput(&got)

		if !reflect.DeepEqual(got, tt.want) {
			// 	// if got := Subsets(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkSubsets(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Subsets(tests[0].arg1)
	}
}

func BenchmarkSubsetsDFS(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubsetsDFS(tests[0].arg1)
	}
}

// sorting [][]int{}
func sortOutput(output *[][]int) {
	for i := range *output {
		sort.Ints((*output)[i])
	}
	sort.Slice(*output, func(i, j int) bool {
		return lexicoOrder(&(*output)[i], &(*output)[j])
	})
}

func lexicoOrder(a, b *[]int) bool {
	for i := 0; i < len(*a) && i < len(*b); i++ {
		if (*a)[i] < (*b)[i] {
			return true
		} else if (*a)[i] > (*b)[i] {
			return false
		}
	}
	return len((*a)) < len((*b))
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=. -cover
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMaxEnvelopes-8          7799131               160.3 ns/op            88 B/op          3 allocs/op
BenchmarkMaxEnvelopes2-8         5800399               195.6 ns/op            80 B/op          4 allocs/op
PASS
coverage: 96.0% of statements
ok      LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes     3.726s
*/
