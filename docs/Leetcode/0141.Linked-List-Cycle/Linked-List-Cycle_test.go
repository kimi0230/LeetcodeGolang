package linkedlistcycle

import "testing"

func TestHasCycle(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}
	// d := &ListNode{4, nil}
	a.Next = b
	b.Next = c
	c.Next = a

	tests := []struct {
		arg1 *ListNode
		want bool
	}{
		{
			arg1: a,
			want: true,
		},
	}

	for _, tt := range tests {
		if got := HasCycle(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
