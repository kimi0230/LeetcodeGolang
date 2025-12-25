package kthlargestelementinanarray

import (
	"container/heap"
	"sort"
)

// 解法一：用Heap資料結構
// container/heap包可以用来構造優先級Queue。
// Heap(堆積)其實是一個Complete Binary Tree(完全二元樹).
// Go的Heap特性是 各個節點都自己是其子樹的根, 且值是最小的(index = 0 的值是最小的).
// 同個根節點的左子樹的值會小於右子樹
func FindKthLargestHeap(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return 0
	}

	// 初始化最小堆
	h := &MinHeap{}
	heap.Init(h)

	// 遍歷集合
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	// fmt.Println(h)

	return (*h)[0]
}

// 自定義最小 Heap 結構體
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	// 小到大排序
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 和 Pop 方法需要使用指針，因為它們會修改 slice 的長度，而不僅僅只內容。
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 解法二：先Sort，再取第k個, 最快！
func FindKthLargestSort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
