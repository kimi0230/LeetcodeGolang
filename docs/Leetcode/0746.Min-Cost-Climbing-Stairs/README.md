---
title: 746. Min Cost Climbing Stairs
subtitle: "https://leetcode.com/problems/min-cost-climbing-stairs/"
date: 2023-12-29T16:24:00+08:00
lastmod: 2023-12-29T16:24:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0746.Min-Cost-Climbing-Stairs"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Easy
  - Min Cost Climbing Stairs
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
# [746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/)

## 題目
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.

 

Example 1:

Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
Example 2:

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.
 

Constraints:

2 <= cost.length <= 1000
0 <= cost[i] <= 999

## 題目大意
陣列的每個索引做為一個階梯，第 i 個階梯對應著一個非負數的體力花費值 cost[i] （索引從 0 開始）。 每當你爬上一個階梯你都要花費對應的體力花費值，然後你可以選擇繼續爬一個階梯或者爬兩個階梯。 您需要找到達到樓層頂部的最低花費。 在開始時，你可以選擇從索引為 0 或 1 的元素作為初始階梯。

## 解題思路
cur 變數存儲從第 i-2 步到達第 i 步的最小花費。
last 變數存儲從第 i-1 步到達第 i 步的最小花費。
在每次迭代中，函數都會比較 cur 和 last 變數的值，並選擇較小的那個存儲在 cur 變數中。
在迭代結束時，last 變數將存儲爬完所有樓梯的最小花費。

cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

| 迭代 | cur | last |
|------|-----|------|
| 2    | 0   | 1    |
| 3    | 1   | 2    |
| 4    | 2   | 2    |
| 5    | 2   | 3    |
| 6    | 3   | 3    |
| 7    | 3   | 4    |
| 8    | 4   | 4    |
| 9    | 4   | 5    |
| 10   | 5   | 6    |


## Big O
時間複雜 : `O(n)`
空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/min-cost-climbing-stairs/
* https://leetcode.cn/problems/min-cost-climbing-stairs/
* https://books.halfrost.com/leetcode/ChapterFour/0700~0799/0746.Min-Cost-Climbing-Stairs/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0746.Min-Cost-Climbing-Stairs/main.go

```go
package mincostclimbingstairs

// 時間複雜 O(n), 空間複雜 O(1)
func MinCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[len(cost)-2], dp[len(cost)-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 時間複雜 O(n), 空間複雜 O(1)
// 優化
// 使用了兩個變數cur和last，
// cur 變數存儲從第 i-2 步到達第 i 步的最小花費。
// last 變數存儲從第 i-1 步到達第 i 步的最小花費。
// 比較兩種選擇的花費：
// 從第2階開始（i := 2），一直迭代到最後一階（i < len(cost)+1）。
// 在每一步驟中：
// 比較兩種選擇的花費：
//
//	選擇從上一階跳到當前階的花費：last + cost[i-1]
//	選擇從上上階跳到當前階的花費：cur + cost[i-2]
//
// 選擇花費較小的方案：
//
//	如果 last + cost[i-1] 更小，則將 last 的值更新為 cur + cost[i-2]，並將 cur 的值更新為 last（即上一步的最小花費）。
//	否則，將 last 的值更新為 last + cost[i-1]，並將 cur 的值更新為 last（即上一步的最小花費）。
func MinCostClimbingStairsOptimize(cost []int) int {
	var cur, last int
	for i := 2; i < len(cost)+1; i++ {
		// 選擇從上一階跳到當前階的花費：last + cost[i-1]
		// 選擇從上上階跳到當前階的花費：cur + cost[i-2]
		if last+cost[i-1] > cur+cost[i-2] {
			cur, last = last, cur+cost[i-2]
		} else {
			cur, last = last, last+cost[i-1]
		}
		// fmt.Printf("%-d | %-d | %-d\n", i, cur, last)
	}
	return last
}

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0746.Min-Cost-Climbing-Stairs
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkMinCostClimbingStairs-4                36693742                30.17 ns/op           24 B/op          1 allocs/op
BenchmarkMinCostClimbingStairsOptimize-4        405489464                3.091 ns/op           0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0746.Min-Cost-Climbing-Stairs   2.713s
```