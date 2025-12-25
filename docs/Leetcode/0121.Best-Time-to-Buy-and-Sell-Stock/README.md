---
title: 0121.Best Time to Buy and Sell Stock
subtitle: "https://leetcode.com/problems/valid-parentheses/"
date: 2023-11-01T17:20:00+08:00
lastmod: 2023-11-01T17:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0121.Best-Time-to-Buy-and-Sell-Stock"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Easy
  - Slide Windows
  - DP
  - Best Time to Buy and Sell Stock
  - array
  - Blind75
  - amazon
  - bloomberg
  - facebook
  - microsoft
  - uber
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
# [0121.Best Time to Buy and Sell Stock]([LEETCODELINK](https://leetcode.com/problems/valid-parentheses/))

## 題目
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

 

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
 

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104
Accepted
3,930,287
Submissions
7,349,138

## 題目大意


## 解題思路

## Big O
時間複雜 : `O(n)`
空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/valid-parentheses/
* https://leetcode.cn/problems/valid-parentheses/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0121.Best-Time-to-Buy-and-Sell-Stock/main.go

```go
package besttimetobuyandsellstock

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func max[T numbers](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// 時間複雜 O(n), 空間複雜 O(1)
// Slide windows
func MaxProfit(prices []int) int {
	left, right := 0, 1
	maxProfit := 0
	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			maxProfit = max(maxProfit, profit)
		} else {
			left = right
		}
		right++
	}
	return maxProfit
}

// 時間複雜 O(n), 空間複雜 O(1)
// DP : 最佳解
func MaxProfitDP(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0] // 初始最低價格為第一天的價格
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// 如果當前價格比最低價格低，更新最低價格
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// 否則計算當前價格賣出時的利潤，並更新最大利潤
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

// 時間複雜 O(), 空間複雜 O()
// DP
func MaxProfitDP2(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		if i-1 == -1 {
			// base case
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0121.Best-Time-to-Buy-and-Sell-Stock
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMaxProfit-8            208734571                8.156 ns/op           0 B/op          0 allocs/op
BenchmarkMaxProfitDP-8          240880467                5.720 ns/op           0 B/op          0 allocs/op
BenchmarkMaxProfitDP2-8          4095866               282.6 ns/op           240 B/op          7 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0121.Best-Time-to-Buy-and-Sell-Stock    6.732s
```