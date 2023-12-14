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

tags: [LeetCode, Go, Easy, Heap, Last Stone Weight]
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


## 解題思路

## Big O
時間複雜 : ``
空間複雜 : ``

## 來源
* https://leetcode.com/problems/last-stone-weight/
* https://leetcode.cn/problems/last-stone-weight/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/1046.Last-Stone-Weight/main.go

```go

```

##  Benchmark

```sh

```