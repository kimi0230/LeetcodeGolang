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