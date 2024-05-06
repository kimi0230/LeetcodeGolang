package middleofthelinkedlist

import (
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	d := &ListNode{4, nil}
	e := &ListNode{5, nil}
	f := &ListNode{6, nil}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = nil

	tests := []struct {
		arg1 *ListNode
		want *ListNode
	}{
		{
			arg1: a,
			want: d,
		},
	}

	for _, tt := range tests {
		if got := MiddleNode(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
