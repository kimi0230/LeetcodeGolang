---
title: 0215. Kth Largest Element in an Array
subtitle: "https://leetcode.com/problems/kth-largest-element-in-an-array/"
date: 2023-07-11T14:33:00+08:00
lastmod: 2023-07-11T14:33:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Kth Largest Element in an Array"
license: ""
images: []

tags: [LeetCode, Go, Medium, Kth Largest Element in an Array, Heap, Priority Queue, Sorting]
categories: [LeetCode]

featuredImage: ""
featuredImagePreview: ""

hiddenFromHomePage: false
hiddenFromSearch: false
twemoji: false
lightgallery: true
ruby: true
fraction: true
fontawesome: true
linkToMarkdown: false
rssFullText: false

toc:
  enable: true
  auto: true
code:
  copy: true
  maxShownLines: 200
math:
  enable: false
  # ...
mapbox:
  # ...
share:
  enable: true
  # ...
comment:
  enable: true
  # ...
library:
  css:
    # someCSS = "some.css"
    # located in "assets/"
    # Or
    # someCSS = "https://cdn.example.com/some.css"
  js:
    # someJS = "some.js"
    # located in "assets/"
    # Or
    # someJS = "https://cdn.example.com/some.js"
seo:
  images: []
  # ...
---
# [0215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

## 題目

Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

 

Example 1:

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
```

Example 2:
```
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
```

Constraints:

* 1 <= k <= nums.length <= 105
* -104 <= nums[i] <= 104

## 題目大意
1.   數組中的第K個最大元素 (Kth Largest Element in an Array)

給定一個未排序的整數數組，找到其中第K個最大的元素。

## 解題思路
一種常見的解法是使用堆數據結構。我們可以維護一個大小為K的最小堆，初始時將數組的前K個元素加入堆中。然後遍歷數組的剩餘元素，如果當前元素比堆頂元素大，則將堆頂元素出堆，並將當前元素加入堆中。最後堆頂元素即為第K個最大的元素。
* 時間複雜度: 構建堆的時間複雜度為O(K)，遍歷數組的時間複雜度為O((N-K)logK)，因此總的時間複雜度為O(NlogK)。
* 空間複雜度: 使用了大小為K的最小堆來存儲元素，因此空間複雜度為O(K)。


在快速選擇 quickselect 的 partition 操作中，每次 partition 操作結束都會返回一個點，這個標定點的下標和最終排序之後有序數組中這個元素所在的下標是一致的。利用這個特性，我們可以不斷的劃分數組區間，最終找到第 K 大的元素。執行一次 partition 操作以後，如果這個元素的下標比 K 小，那麼接著就在後邊的區間繼續執行 partition 操作；如果這個元素的下標比 K 大，那麼就在左邊的區間繼續執行 partition 操作；如果相等就直接輸出這個下標對應的數組元素即可。
快速選擇 quickselect 的思路實現的算法時間複雜度為 O(n)，空間複雜度為 O(logn)

## 來源
* https://leetcode.com/problems/kth-largest-element-in-an-array/
* https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0215.Kth-Largest-Element-in-an-Array/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0215.Kth-Largest-Element-in-an-Array/main.go

```go
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

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0215.Kth-Largest-Element-in-an-Array
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkFindKthLargestHeap-8            6112726               184.9 ns/op            48 B/op          3 allocs/op
BenchmarkFindKthLargestSort-8           18023403                60.38 ns/op           24 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0215.Kth-Largest-Element-in-an-Array    3.383s
```