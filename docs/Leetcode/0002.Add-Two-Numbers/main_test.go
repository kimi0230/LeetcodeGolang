package addtwonumbers

import (
	"reflect"
	"testing"
)

// Ints2List converts []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// List2Ints converts List to []int
func List2Ints(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		l1 := Ints2List(tt.l1)
		l2 := Ints2List(tt.l2)
		got := addTwoNumbers(l1, l2)
		gotArr := List2Ints(got)

		if !reflect.DeepEqual(gotArr, tt.want) {
			t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, gotArr, tt.want)
		}
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	l1 := Ints2List([]int{2, 4, 3})
	l2 := Ints2List([]int{5, 6, 4})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}
