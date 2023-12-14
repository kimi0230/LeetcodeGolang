package laststoneweight

import "testing"

var tests = []struct {
	arg1 []int
	want int
}{
	{
		[]int{2, 7, 4, 1, 8, 1},
		1,
	},
}

func TestLastStoneWeight(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := LastStoneWeight(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkLastStoneWeight(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastStoneWeight(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/1046.Last-Stone-Weight -bench=.

*/
