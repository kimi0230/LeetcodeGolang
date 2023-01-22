package reverselinkedlist

import (
	. "LeetcodeGolang/structures"
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 *ListNode
	want *ListNode
}{
	{
		Ints2List([]int{1, 2, 3, 4, 5}),
		Ints2List([]int{5, 4, 3, 2, 1}),
	},
	{
		Ints2List([]int{1, 2}),
		Ints2List([]int{2, 1}),
	},
	{
		Ints2List([]int{}),
		Ints2List([]int{}),
	},
}

func TestReverseList(t *testing.T) {
	for _, tt := range tests {
		if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkReverseList(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseList(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0206.Reverse-Linked-List -bench=.

*/
