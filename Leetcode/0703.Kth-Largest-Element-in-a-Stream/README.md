---
title: 0703. Kth Largest Element in a Stream
subtitle: "https://leetcode.com/problems/kth-largest-element-in-a-stream/"
date: 2023-01-22T21:20:00+08:00
lastmod: 2023-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0703.Kth-Largest-Element-in-a-Stream"
license: ""
images: []

tags: [LeetCode, Go, Easy, Kth Largest Element in a Stream, Heap, Priority Queue]
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
# [0703. Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/)

## 題目
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.
 

Example 1:

Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]

Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
 

Constraints:

1 <= k <= 104
0 <= nums.length <= 104
-104 <= nums[i] <= 104
-104 <= val <= 104
At most 104 calls will be made to add.
It is guaranteed that there will be at least k elements in the array when you search for the kth element.
Accepted
479.1K
Submissions
846.4K
Acceptance Rate
56.6%

## 題目大意
設計一個找到數據流中第 k 大元素的類（class）。 注意是排序後的第 k 大元素，不是第 k 個不同的元素。 請實現 KthLargest 類：
* KthLargest（int k， int[] nums） 使用整數 k 和整數流 nums 初始化物件。
* int add（int val） 將 val 插入數據流 nums 後，返回當前數據流中第 k 大的元素。

## 解題思路
這題考察優先順序佇列的使用，可以先做下這道類似的題目 [215.陣列中的第 K 個最大元素](../0215.Kth-Largest-Element-in-an-Array/README.md)。

[golang container/heap](../../structures/heap/index.en.md)

## Big O
時間複雜 : ``
空間複雜 : ``

## 來源
* https://leetcode.com/problems/kth-largest-element-in-a-stream/
* https://leetcode.cn/problems/kth-largest-element-in-a-stream/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0703.Kth-Largest-Element-in-a-Stream/main.go

```go

```

##  Benchmark

```sh

```