package kthlargestelementinastream

import (
	"container/heap"
	"sort"
)

// 時間複雜 :
// 初始化時間複雜度為： O(nlog⁡k) ，其中 n 為初始化時 nums 的長度;
// 單次插入時間複雜度為： O(log⁡k)

// 空間複雜 : O(k)。 需要使用優先佇列存儲前 k 大的元素
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

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func KthLargestStream(k int, nums []int, elem []int) []int {
	obj := Constructor(k, nums)
	result := []int{0}
	for _, val := range elem {
		obj.Add(val)
		result = append(result, obj.IntSlice[0])
	}

	return result
}
