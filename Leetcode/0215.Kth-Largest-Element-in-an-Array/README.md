---
title: 0215. Kth Largest Element in an Array
subtitle: "https://leetcode.com/problems/kth-largest-element-in-an-array/"
date: 2023-07-11T14:33:00+08:00
lastmod: 2023-07-11T14:33:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Kth Largest Element in an Array"
license: ""
images: []

tags: [LeetCode, Go, Medium, Kth Largest Element in an Array, Heap, Sorting]
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
# [0215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

## 題目

Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

 

Example 1:

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
```

Example 2:
```
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
```

Constraints:

* 1 <= k <= nums.length <= 105
* -104 <= nums[i] <= 104

## 題目大意
1.   數組中的第K個最大元素 (Kth Largest Element in an Array)

給定一個未排序的整數數組，找到其中第K個最大的元素。

## 解題思路
一種常見的解法是使用堆數據結構。我們可以維護一個大小為K的最小堆，初始時將數組的前K個元素加入堆中。然後遍歷數組的剩餘元素，如果當前元素比堆頂元素大，則將堆頂元素出堆，並將當前元素加入堆中。最後堆頂元素即為第K個最大的元素。
* 時間複雜度: 構建堆的時間複雜度為O(K)，遍歷數組的時間複雜度為O((N-K)logK)，因此總的時間複雜度為O(NlogK)。
* 空間複雜度: 使用了大小為K的最小堆來存儲元素，因此空間複雜度為O(K)。

## 來源
* https://leetcode.com/problems/kth-largest-element-in-an-array/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0215.Kth-Largest-Element-in-an-Array/main.go

```go

```

##  Benchmark

```sh

```