package topkfrequentelements

import (
	"container/heap"
	"sort"
)

// 方法一: 使用 PriorityQueue
// 時間複雜 O(Nlog⁡k), 空間複雜 O(N)
func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	q := PriorityQueue{}
	for key, count := range m {
		item := &Item{key: key, count: count}
		q.PushPQ(item)
	}

	var result []int
	for len(result) < k {
		item := q.PopPQ()
		result = append(result, item.key)
	}
	return result
}

// Item define
type Item struct {
	key   int
	count int
}

/*
// A PriorityQueue implements heap.Interface and holds Items.
package heap

import "sort"

type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}
*/

type PriorityQueue []*Item

/*
// An implementation of Interface can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.

	type Interface interface {
		// Len is the number of elements in the collection.
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}
*/
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 大到小排序
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pg *PriorityQueue) Push(v interface{}) {
	item := v.(*Item)
	*pg = append(*pg, item)
}

func (pg *PriorityQueue) Pop() interface{} {
	n := len(*pg)
	item := (*pg)[n-1]
	*pg = (*pg)[:n-1]
	return item
}

func (pg *PriorityQueue) PushPQ(v *Item) {
	heap.Push(pg, v)
}

func (pg *PriorityQueue) PopPQ() *Item {
	return heap.Pop(pg).(*Item)
}

// 方法二: 使用 Quick Sort
// 時間複雜 O(), 空間複雜 O()
func TopKFrequentQuickSort(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	values := [][]int{}
	for key, count := range m {
		values = append(values, []int{key, count})
	}

	result := []int{}
	sort.Sort(sortValue(values))

	for i := 0; i < k; i++ {
		result = append(result, values[i][0])
	}
	return result
}

type sortValue [][]int

func (s sortValue) Len() int {
	return len(s)
}

func (s sortValue) Less(i, j int) bool {
	return s[i][1] > s[j][1]
}

func (s sortValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
