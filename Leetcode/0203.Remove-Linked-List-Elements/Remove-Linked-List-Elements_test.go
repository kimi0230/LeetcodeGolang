package removelinkedlistelements

import (
	"LeetcodeGolang/structures"
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 *ListNode
	arg2 int
	want *ListNode
}{
	{
		structures.Ints2List([]int{1, 2, 6, 3, 4, 5, 6}),
		6,
		structures.Ints2List([]int{1, 2, 3, 4, 5}),
	},
	{
		structures.Ints2List([]int{7, 7, 7, 7, 7}),
		7,
		structures.Ints2List([]int{}),
	},
	{
		structures.Ints2List([]int{1, 2, 3, 4, 5}),
		2,
		structures.Ints2List([]int{1, 3, 4, 5}),
	},
	{
		structures.Ints2List([]int{1, 2, 3, 2, 3, 2, 3, 2}),
		2,
		structures.Ints2List([]int{1, 3, 3, 3}),
	},
	{
		structures.Ints2List([]int{1, 2, 3, 4, 5}),
		10,
		structures.Ints2List([]int{1, 2, 3, 4, 5}),
	},
	{
		structures.Ints2List([]int{1}),
		1,
		structures.Ints2List([]int{}),
	},
}

func TestRemoveElements(t *testing.T) {
	for _, tt := range tests {
		if got := RemoveElements(tt.arg1, tt.arg2); !reflect.DeepEqual(structures.List2Ints(got), structures.List2Ints(tt.want)) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
		// got := RemoveElements(tt.arg1, tt.arg2)
		// fmt.Println(structures.List2Ints(got))
	}
}

func TestRemoveElementsRecurse(t *testing.T) {
	for _, tt := range tests {
		if got := RemoveElementsRecurse(tt.arg1, tt.arg2); !reflect.DeepEqual(structures.List2Ints(got), structures.List2Ints(tt.want)) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
		// got := RemoveElements(tt.arg1, tt.arg2)
		// fmt.Println(structures.List2Ints(got))
	}
}

func TestRemoveElementsVirtualNode(t *testing.T) {
	for _, tt := range tests {
		if got := RemoveElementsVirtualNode(tt.arg1, tt.arg2); !reflect.DeepEqual(structures.List2Ints(got), structures.List2Ints(tt.want)) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkRemoveElements(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElements(tests[i%6].arg1, tests[i%6].arg2)
	}
}
func BenchmarkRemoveElementsVirtualNode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElementsVirtualNode(tests[i%6].arg1, tests[i%6].arg2)
	}
}

func BenchmarkRemoveElementsRecurse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveElementsRecurse(tests[i%6].arg1, tests[i%6].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0203.Remove-Linked-List-Elements -bench=.
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkRemoveElements-8               195664567                5.953 ns/op           0 B/op          0 allocs/op
BenchmarkRemoveElementsVirtualNode-8    217690257                4.806 ns/op           0 B/op          0 allocs/op
BenchmarkRemoveElementsRecurse-8        79181312                14.31 ns/op            0 B/op          0 allocs/op
*/
