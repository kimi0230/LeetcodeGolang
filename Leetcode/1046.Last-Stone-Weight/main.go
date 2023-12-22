package laststoneweight

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x IntSlice) Sort() { Sort(x) }
*/

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	// 大到小排序
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	old := h.IntSlice
	v := old[len(old)-1]
	h.IntSlice = old[:len(old)-1]
	return v
}

func (h *hp) PushInt(v int) {
	heap.Push(h, v)
}

func (h *hp) PopInt() int {
	return heap.Pop(h).(int)
}

// 時間複雜 O(nlogn), 空間複雜 O(n)
// n 是石頭數量. 每次從隊列中取出元素需要話O(logn) 的時間, 最多共需要需要粉碎 n−1 次石頭
func LastStoneWeight(stones []int) int {
	q := &hp{stones}
	heap.Init(q)
	fmt.Println(q)
	for q.Len() > 1 {
		fmt.Println(q)
		x, y := q.PopInt(), q.PopInt()
		fmt.Printf("%d,%d\n", x, y)
		if x > y {
			q.PushInt(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}
