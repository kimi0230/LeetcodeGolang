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


## 解題思路

* 簡單的 DP，經典的爬樓梯問題. 一個樓梯可以由 n-1 和 n-2 的樓梯爬上來。
* 這一題求解的值就是斐波那契數列。

## Big O
時間複雜 : ``
空間複雜 : ``

## 來源
* https://leetcode.com/problems/climbing-stairs/
* https://leetcode.cn/problems/climbing-stairs/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0070.Climbing-Stairs/main.go

```go

```

##  Benchmark

```sh

```