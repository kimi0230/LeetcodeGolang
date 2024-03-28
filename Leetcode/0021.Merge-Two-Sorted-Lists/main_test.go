package mergetwosortedlists

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []int
	arg2 []int
	want []int
}{
	{
		[]int{1, 2, 4},
		[]int{1, 3, 4},
		[]int{1, 1, 2, 3, 4, 4},
	},
}

func TestMergeTwoLists(t *testing.T) {
	for _, tt := range tests {
		arg1 := &ListNode{}
		arg1Cur := arg1
		for _, v := range tt.arg1 {
			tmp := &ListNode{Val: v}
			arg1Cur.Next = tmp
			arg1Cur = arg1Cur.Next
		}

		arg2 := &ListNode{}
		arg2Cur := arg2
		for _, v := range tt.arg2 {
			tmp := &ListNode{Val: v}
			arg2Cur.Next = tmp
			arg2Cur = arg2Cur.Next
		}

		wantNode := &ListNode{}
		wantNodeCur := wantNode
		for _, v := range tt.want {
			tmp := &ListNode{Val: v}
			wantNodeCur.Next = tmp
			wantNodeCur = wantNodeCur.Next
		}

		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := MergeTwoLists(arg1.Next, arg2.Next); !reflect.DeepEqual(got, wantNode.Next) {
			t.Errorf("got = %v, want = %v", got, wantNode.Next)
		}
	}
}

func BenchmarkMergeTwoLists(b *testing.B) {
	b.ResetTimer()
	arg1 := &ListNode{}
	arg1Cur := arg1
	for _, v := range tests[0].arg1 {
		tmp := &ListNode{Val: v}
		arg1Cur.Next = tmp
		arg1Cur = arg1Cur.Next
	}

	arg2 := &ListNode{}
	arg2Cur := arg2
	for _, v := range tests[0].arg2 {
		tmp := &ListNode{Val: v}
		arg2Cur.Next = tmp
		arg2Cur = arg2Cur.Next
	}
	for i := 0; i < b.N; i++ {
		MergeTwoLists(arg1.Next, arg2.Next)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
