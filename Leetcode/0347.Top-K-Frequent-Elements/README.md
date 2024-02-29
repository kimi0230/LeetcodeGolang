---
title: 347. Top K Frequent Elements
subtitle: "https://leetcode.com/problems/top-k-frequent-elements/description/"
date: 2024-01-17T17:12:00+08:00
lastmod: 2024-01-17T17:12:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0347.Top-K-Frequent-Elements"
license: ""
images: []

tags: [LeetCode, Go, Medium, Top K Frequent Elements, heap]
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
# [347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/description/)

## 題目
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

 

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
 

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.
 

Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

Seen this question in a real interview before? 1/4

* Accepted 1.9M
* Submissions 3M
* Acceptance Rate 62.7%

## 題目大意


## 解題思路
### 方法一：[heap](../../README.md#heap--priority-queue)
方法一：堆
思路與演算法

首先遍歷整個數組，並使用哈希表記錄每個數字出現的次數，並形成一個「出現次數數組」。 找出原數組的前 k 個高頻元素，就等於找出「出現次數數組」的前 k 大的值。

最簡單的做法是為「出現次數數組」排序。 但由於可能有O(N) 個不同的出現次數（其中N 為原數組長度），故總的演算法複雜度會達到O(Nlog⁡N)，不符合題目的要求。

在這裡，我們可以利用堆的想法：建立一個小頂堆，然後遍歷「出現次數數組」：

如果堆的元素個數小於 k，就可以直接插入堆中。
如果堆的元素個數等於 k，則檢查堆頂與目前出現次數的大小。 如果堆頂較大，表示至少有 k 個數字的出現次數比目前值大，故捨棄目前值；否則，就彈出堆頂，並將目前值插入堆中。
遍歷完成後，堆中的元素就代表了「出現次數數組」中前 k 大的值。

複雜度分析

時間複雜度：`O(Nlog⁡k)`，
其中 N 為陣列的長度。 我們先遍歷原數組，並使用雜湊表記錄出現次數，每個元素需要O(1) 的時間，共需`O(N)` 的時間 。 
隨後，我們遍歷「出現次數數組」，由於堆的大小至多為k，因此每次堆操作需要`O(log⁡k)`的時間，共需`O(Nlog⁡k)`的時間。 二者之和為 `O(Nlog⁡k)`。

空間複雜度：`O(N)`。 
雜湊表的大小為`O(N)`，而堆的大小為`O(k)`，共為`O(N)`。

### 方法二： Quick Sort
| Name       | Best     | Average  | Worst | Memory | Stable |
|------------|----------|----------|-------|--------|--------|
| Quick sort | n log(n) | n log(n) | n^2   | log(n) | No     |

使用基於快速排序的方法，求出「出現次數陣列」的前 k 大的值。


## Big O

### 方法一：heap
時間複雜 : `O(Nlog⁡k)`
空間複雜 : `O(N)`

### 方法二： Quick Sort
時間複雜 : ``
空間複雜 : ``

## 來源
* https://leetcode.com/problems/top-k-frequent-elements/description/
* https://leetcode.cn/problems/top-k-frequent-elements/description/
* https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0347.Top-K-Frequent-Elements/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0347.Top-K-Frequent-Elements/main.go

```go
package topkfrequentelements

import (
	"container/heap"
	"sort"
)

// 方法一: 使用 PriorityQueue
// 時間複雜 O(Nlog⁡k), 空間複雜 O(N)
// 首先遍歷整個數組，並使用哈希表記錄每個數字出現的次數，並形成一個「出現次數數組」
// 建立一個 PriortyQueue, 將「出現次數數組」丟進去
// 在把 PriortyQueue pop的值丟到 result
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

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0347.Top-K-Frequent-Elements
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkTopKFrequent-8                  1875834               648.6 ns/op           200 B/op         11 allocs/op
BenchmarkTopKFrequentQuickSort-8         1830498               704.2 ns/op           312 B/op         11 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0347.Top-K-Frequent-Elements    3.831s
```