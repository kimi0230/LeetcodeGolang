package oddoccurrencesinarray

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{9, 3, 9, 3, 9, 7, 9},
		7,
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range tests {
		if got := Solution(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, tt := range tests {
		if got := Solution2(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution(tests[0].arg1)
	}
}
func BenchmarkSolution2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution2(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Codility/Lesson/0002.Array/OddOccurrencesInArray -bench=.
BenchmarkSolution-4      5143094               224.6 ns/op            48 B/op          2 allocs/op
BenchmarkSolution2-4    335248345                3.501 ns/op           0 B/op          0 allocs/op
*/
