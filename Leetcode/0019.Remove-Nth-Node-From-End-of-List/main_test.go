package removenthnodefromendoflist

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = nil

	tests := []struct {
		arg1 *ListNode
		arg2 int
		want []int
	}{
		{
			arg1: a,
			arg2: 2,
			want: []int{1, 2, 3, 5},
		},
	}

	for _, tt := range tests {
		got := RemoveNthFromEnd(tt.arg1, tt.arg2)

		gotArr := []int{}
		for got != nil {
			gotArr = append(gotArr, got.Val)
			got = got.Next
		}
		if !reflect.DeepEqual(gotArr, tt.want) {
			t.Errorf("got = %v, want = %v", gotArr, tt.want)
		}
	}
}
