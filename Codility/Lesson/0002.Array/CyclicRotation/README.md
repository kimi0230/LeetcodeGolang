---
title: Cyclic Rotation
subtitle: "https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/"
date: 2021-04-18T13:10:00+08:00
lastmod: 2021-04-18T13:10:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Cyclic Rotation"
license: ""
images: []

tags: [Codility, Go, Array, Painless, Multiple Pointers]
categories: [Codility]

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

# [CyclicRotation](https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/)

An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).

The goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.

Write a function:

func Solution(A []int, K int) []int

that, given an array A consisting of N integers and an integer K, returns the array A rotated K times.

For example, given

    A = [3, 8, 9, 7, 6]
    K = 3
the function should return [9, 7, 6, 3, 8]. Three rotations were made:

    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]
For another example, given

    A = [0, 0, 0]
    K = 1
the function should return [0, 0, 0]

Given

    A = [1, 2, 3, 4]
    K = 4
the function should return [1, 2, 3, 4]

Assume that:

N and K are integers within the range [0..100];
each element of array A is an integer within the range [−1,000..1,000].
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
CyclicRotation題目要求將給定的整數陣列按照給定的旋轉步數進行循環右移，並返回旋轉後的陣列。例如，如果陣列是[3, 8, 9, 7, 6]且旋轉步數是3，則右移後的陣列為[9, 7, 6, 3, 8]。

## 解題思路
解題思路可以使用多種方法。一種常見的方法是使用額外的陣列來存儲旋轉後的結果。另一種方法是通過循環右移的操作，直接在原始陣列上進行元素交換。根據旋轉步數，我們可以將陣列分為兩個部分，並進行相應的元素交換操作。

時間複雜度: 解題思路中的操作需要遍歷整個陣列，因此時間複雜度為O(N)，其中N是陣列的長度。
空間複雜度: 解題思路中使用了額外的陣列或進行原地交換，不需要使用額外的數據結構，因此空間複雜度為O(1)。

## 來源
* https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0002.Array/CyclicRotation/CyclicRotation.go


```go
package cyclicrotation

func Solution(A []int, K int) []int {
	if K == 0 || len(A) <= 1 {
		return A
	}

	K = K % len(A)
	return append(A[len(A)-K:], A[:len(A)-K]...)
}

func Solution2(A []int, K int) []int {
	if K == 0 || len(A) <= 1 {
		return A
	}
	if K > len(A) {
		K = K % len(A)
	}
	return append(A[len(A)-K:], A[:len(A)-K]...)
}
```