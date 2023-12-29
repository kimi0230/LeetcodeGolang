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

tags: [LeetCode, Go, Easy, Min Cost Climbing Stairs]
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


## 解題思路
cur 變數存儲從第 i-2 步到達第 i 步的最小花費。
last 變數存儲從第 i-1 步到達第 i 步的最小花費。
在每次迭代中，函數都會比較 cur 和 last 變數的值，並選擇較小的那個存儲在 cur 變數中。
在迭代結束時，last 變數將存儲爬完所有樓梯的最小花費。

cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

| 迭代 | i  | cur | last |
|------|----|-----|------|
| 1    | 2  | 100 | 1    |
| 2    | 3  | 101 | 100  |
| 3    | 4  | 101 | 101  |
| 4    | 5  | 102 | 101  |
| 5    | 6  | 202 | 102  |
| 6    | 7  | 203 | 202  |
| 7    | 8  | 303 | 203  |
| 8    | 9  | 304 | 303  |
| 9    | 10 | 404 | 304  |
| 10   | 11 | 405 | 404  |

## Big O
時間複雜 : ``
空間複雜 : ``

## 來源
* https://leetcode.com/problems/min-cost-climbing-stairs/
* https://leetcode.cn/problems/min-cost-climbing-stairs/
* https://books.halfrost.com/leetcode/ChapterFour/0700~0799/0746.Min-Cost-Climbing-Stairs/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0746.Min-Cost-Climbing-Stairs/main.go

```go

```

##  Benchmark

```sh

```