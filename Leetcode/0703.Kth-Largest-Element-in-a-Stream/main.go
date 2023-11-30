package kthlargestelementinastream

import (
	"container/heap"
	"sort"
)

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// 時間複雜 O(), 空間複雜 O()
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	// 拿出最後一筆
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func KthLargestStream(k int, nums []int, elem []int) []int {
	obj := Constructor(k, nums)
	result := []int{0}
	for _, val := range elem {
		obj.Add(val)
		result = append(result, obj.IntSlice[0])
	}

	return result
}
