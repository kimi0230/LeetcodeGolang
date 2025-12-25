package linkedlistcycleII

import (
	"reflect"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	// d := &ListNode{4, nil}
	a.Next = b
	b.Next = c
	c.Next = a

	d := &ListNode{1, nil}
	e := &ListNode{2, nil}
	f := &ListNode{3, nil}
	d.Next = e
	e.Next = f
	f.Next = a

	tests := []struct {
		arg1 *ListNode
		want *ListNode
	}{
		{
			arg1: a,
			want: a,
		},
		{
			arg1: d,
			want: a,
		},
	}

	for _, tt := range tests {
		if got := DetectCycle(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
