---
title: 1046. Last Stone Weight
subtitle: "https://leetcode.com/problems/last-stone-weight/"
date: 2023-12-14T17:38:00+08:00
lastmod: 2023-12-14T17:38:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "1046.Last-Stone-Weight"
license: ""
images: []

tags: [LeetCode, Go, Easy, Heap, Priority Queue, Last Stone Weight]
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
# [1046. Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)

## 題目
You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:

If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.

 

Example 1:

Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation: 
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
Example 2:

Input: stones = [1]
Output: 1
 

Constraints:

1 <= stones.length <= 30
1 <= stones[i] <= 1000

## 題目大意
有一個集合 stones，每個 stone 的重量由正整數表示。
每次可以選擇兩個不同的石頭，將它們一起粉碎，然後得到一個新的石頭，其重量為兩者之差。
你需要重複這個過程，直到集合中只剩下一個石頭，或者集合中沒有石頭為止。
在這個過程中，找到可能的最後一顆石頭的重量。如果集合中沒有石頭，則返回 0。

```makefile
Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
步驟1：選擇石頭 7 和 8，得到新石頭 [2,4,1,1,1]。
步驟2：選擇石頭 2 和 4，得到新石頭 [2,1,1,1]。
步驟3：選擇石頭 2 和 1，得到新石頭 [1,1,1]。
步驟4：選擇石頭 1 和 1，得到新石頭 [0,1]。
步驟5：選擇石頭 1 和 0，得到新石頭 [1]。
最後剩下的石頭的重量為 1。
```

## 解題思路

1. 將 stones 陣列轉換為最大堆（max heap），可以使用優先佇列實現。
2. 進行迴圈，每次從最大堆中取出兩個最大的石頭。
3. 如果兩個石頭不相等，將它們的差值插入最大堆。
4. 重複上述步驟，直到最大堆中只剩下一個石頭或沒有石頭為止。
5. 如果最大堆中有石頭，返回該石頭的重量，否則返回 0。
這樣的做法確保每次都選擇最大的兩個石頭進行粉碎，最終留下的石頭重量就是可能的最後一個石頭的重量。

參考 [0215 Kth Largest Element in an Array](../0215.Kth-Largest-Element-in-an-Array/README.md)

## Big O
時間複雜 : `O(nlogn)`
空間複雜 : `O(n)`

`n` 是石頭數量. 每次從隊列中取出元素需要話`O(logn)` 的時間, 最多共需要需要粉碎 `n−1` 次石頭

## 來源
* https://leetcode.com/problems/last-stone-weight/
* https://leetcode.cn/problems/last-stone-weight/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/1046.Last-Stone-Weight/main.go

```go
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

```

##  Benchmark

```sh

```