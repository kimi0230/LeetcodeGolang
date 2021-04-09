package twosum

import (
	"reflect"
	"testing"
)

func TestTwosum(t *testing.T) {
	tests := []struct {
		name string
		arg1 []int
		arg2 int
		want []int
	}{
		{
			name: "Twosum",
			arg1: []int{2, 7, 11, 15},
			arg2: 9,
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Twosum(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestTwosum2(t *testing.T) {
	tests := []struct {
		name string
		arg1 []int
		arg2 int
		want []int
	}{
		{
			name: "Twosum2",
			arg1: []int{2, 7, 11, 15},
			arg2: 9,
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Twosum2(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestTwosum3(t *testing.T) {
	tests := []struct {
		name string
		arg1 []int
		arg2 int
		want []int
	}{
		{
			name: "Twosum3",
			arg1: []int{2, 7, 11, 15},
			arg2: 9,
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Twosum3(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTwosum(b *testing.B) {
	arg1 := []int{2, 7, 11, 15}
	arg2 := 9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Twosum(arg1, arg2)
	}
}

func BenchmarkTwosum2(b *testing.B) {
	arg1 := []int{2, 7, 11, 15}
	arg2 := 9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Twosum2(arg1, arg2)
	}
}

func BenchmarkTwosum3(b *testing.B) {
	arg1 := []int{2, 7, 11, 15}
	arg2 := 9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Twosum3(arg1, arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/0001.Two-Sum -bench=.
BenchmarkTwosum-8       53737875                22.04 ns/op           16 B/op          1 allocs/op
BenchmarkTwosum2-8      25733203                44.74 ns/op           16 B/op          1 allocs/op
BenchmarkTwosum3-8      43901755                26.22 ns/op           16 B/op          1 allocs/op
*/
