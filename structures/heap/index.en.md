---
title: "Golang: container/heap"
subtitle: "https://leetcode.com/problems/kth-largest-element-in-a-stream/"
date: 2023-01-22T21:20:00+08:00
lastmod: 2023-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0703.Kth-Largest-Element-in-a-Stream"
license: ""
images: []

tags: [Go, container/heap]
categories: [Structure]

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

# Golang : container/heap

## Heap是什麼
Wiki: https://zh.wikipedia.org/wiki/%E5%A0%86%E7%A9%8D

堆積（Heap）是電腦科學中的一種特別的**完全二元樹**。
若是滿足以下特性，即可稱為堆積：「給定堆積中任意節點P和C，若P是C的母節點，那麼P的值會小於等於（或大於等於）C的值」。
若母節點的值恆小於等於子節點的值，此堆積稱為最小堆積（min heap）；
反之，若母節點的值恆大於等於子節點的值，此堆積稱為最大堆積（max heap）。
在堆積中最頂端的那一個節點，稱作根節點（root node），根節點本身沒有母節點（parent node）。

![](FullBT_CompleteBT.jpg)

## container/heap 提供的方法
heap包為實現了 `heap.Interface` 的類型提供了堆方法：Init/Push/Pop/Remove/Fix。 `container/heap` 為最小堆，
即每個節點的值都小於它的子樹的所有元素的值（A heap is a tree with the property that each node is the minimum-valued node in its subtree）。

**heap:**
```go
package heap
// ...
// Note that Push and Pop in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use heap.Push and heap.Pop.
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}
```

**sort:** 
```go
package sort

// An implementation of Interface can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

由於 heap.Interface 包含了 sort.Interface ，所以，目標類型需要包含如下方法：`Len`/`Less`/`Swap` ,`Push`/`Pop。`

## container/heap 可用在哪

container/heap包可以用來構造優先順序佇列。

https://go.dev/play/p/77zrF3PurO4
```go
// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
```

PriorityQueue 本質上是個 *Item 陣列，其Len/Less/Swap是比較常見的陣列用來sort需要定義的函數，而Push、Pop則是使用數位來插入、的方法。 PriorityQueue 還提供了update方法。 注意由於通常希望優先順序佇列Pop出來的是優先順序最高的元素，所以Less方法是反著寫的。

定義了以上方法以後， PriorityQueue 就具備了使用 container/heap 包的條件。


如下代碼，先從items map出發定義了一個pq陣列，長度為hash的size，並調用 heap.Init 初始化pq; 
之後向佇列中增加了一個優先順序為1的元素，並更新該元素的佇列; 最後從佇列中依此Pop，可見元素在Pop時是依照優先順序排序的。
```go
// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s index:%d \n", item.priority, item.value, item.index)
	}
}
// Output:
// 05:orange index:-1 
// 04:pear index:-1 
// 03:banana index:-1 
// 02:apple index:-1 

```

## Reference
* [Golang: 详解container/heap](https://ieevee.com/tech/2018/01/29/go-heap.html#3-containerheap%E5%8F%AF%E4%BB%A5%E7%94%A8%E6%9D%A5%E5%81%9A%E4%BB%80%E4%B9%88)