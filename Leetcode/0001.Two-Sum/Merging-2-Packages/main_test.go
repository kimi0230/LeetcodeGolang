package merging2packages

import (
	"reflect"
	"testing"
)

func TestGetIndicesOfItemWeights(t *testing.T) {
	tests := []struct {
		name string
		arg1 []int
		arg2 int
		want []int
	}{
		{
			name: "getIndicesOfItemWeights",
			arg1: []int{4, 6, 10, 15, 16},
			arg2: 21,
			want: []int{3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIndicesOfItemWeights(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetIndicesOfItemWeights(b *testing.B) {
	arg1 := []int{2, 7, 11, 15}
	arg2 := 9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIndicesOfItemWeights(arg1, arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0001.Two-Sum -bench=.
BenchmarkTwosum-8       53737875                22.04 ns/op           16 B/op          1 allocs/op
BenchmarkTwosum2-8      25733203                44.74 ns/op           16 B/op          1 allocs/op
*/
