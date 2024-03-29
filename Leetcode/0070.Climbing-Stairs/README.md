---
title: 0070.Climbing Stairs
subtitle: "https://leetcode.com/problems/climbing-stairs/"
date: 2023-12-27T16:34:00+08:00
lastmod: 2023-12-27T16:34:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0070.Climbing-Stairs"
license: ""
images: []

tags: [LeetCode, Go, Easy, Climbing Stairs, DP]
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
# [0070.Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

## 題目
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

 

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 

Constraints:

1 <= n <= 45

* Accepted: 2.8M
* Submissions: 5.4M
* Acceptance Rate: 52.3%

## 題目大意
類似 [Fibonacci Number](../0509.Fibonacci-Number/README.md)

## 解題思路

* 簡單的 DP，經典的爬樓梯問題. 一個樓梯可以由 n-1 和 n-2 的樓梯爬上來。
* 這一題求解的值就是斐波那契數列。

## Big O
時間複雜 : ``
空間複雜 : ``

## 來源
* https://leetcode.com/problems/climbing-stairs/
* https://leetcode.cn/problems/climbing-stairs/
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0070.Climbing-Stairs/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0070.Climbing-Stairs/main.go

```go

package climbingstairs

// 時間複雜 O(n), 空間複雜 O(n)
func ClimbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func ClimbStairsCache(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		if val := dp[i]; val == 0 {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}

func ClimbStairsRecursive(n int) int {
	dp := make([]int, n+1)
	// dp[0], dp[1] = 1, 1

	var climbClosure func(n int) int
	climbClosure = func(n int) int {
		if n <= 2 {
			return n
		}
		dp[n] = climbClosure(n-1) + climbClosure(n-2)
		return dp[n]
	}
	return climbClosure(n)
}
```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0070.Climbing-Stairs
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkClimbStairs-8                  10386211               112.1 ns/op           320 B/op          1 allocs/op
BenchmarkClimbStairsCache-8             10184984               118.8 ns/op           320 B/op          1 allocs/op
BenchmarkClimbStairsRecursive-8                4         281980486 ns/op             320 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0070.Climbing-Stairs    5.591s
```