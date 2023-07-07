---
title: Binary Gap
subtitle: "https://app.codility.com/programmers/lessons/1-iterations/binary_gap/"
date: 2021-04-18T12:09:00+08:00
lastmod: 2021-04-18T12:09:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Binary Gap"
license: ""
images: []

tags: [Codility, Go, Iterations, Painless, Bitwise Manipulation]
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

# [BinaryGap](https://app.codility.com/programmers/lessons/1-iterations/binary_gap/)
START
Find longest sequence of zeros in binary representation of an integer.

## 題目
A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

func Solution(N int) int

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
輸入正整數, 找出此數在二進位,兩個bit值為1中裡面隔著最多0的長度

## 解題思路
先找出bit 1的位子, 再開始算中間最長的長度
解題思路通常可以使用位運算來處理二進制數字。我們可以將N轉換為二進制表示，然後使用遍歷或迴圈來找到相鄰1之間的最大距離。可以使用兩個指針來記錄相鄰的1的

時間複雜度: 解題思路中的遍歷或迴圈需要將N轉換為二進制，因此時間複雜度取決於二進制表示的位數。假設N的位數為k，則時間複雜度為O(k)。

空間複雜度: 解題思路中不需要使用額外的數據結構，只需要使用幾個變數來保存位置和計算結果，因此空間複雜度為O(1)。

## 來源
* https://app.codility.com/programmers/lessons/1-iterations/binary_gap/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0001.Iterations/Binary-Gap/Binary-Gap.go


```go
package binarygap

// O(log n)
func Solution(N int) int {
	maxLen, curLen := 0, 0
	findOne := false
	for N > 0 {
		curBit := N & 1
		if curBit == 1 {
			curLen = 0
			findOne = true
		} else if curBit == 0 && findOne {
			curLen++
		}

		if curLen > maxLen {
			maxLen = curLen
		}
		N = N >> 1
	}
	return maxLen
}

// https://wandbox.org/permlink/totZwDAbL1wCgsqt
func evil(x int) int {
	if x&(x+1) > 0 {
		return evil(x|(x>>1)) + 1
	} else {
		return 0
	}
}

func SolutionRecur(N int) int {
	for (N & 1) == 0 {
		N = N >> 1
	}
	return evil(N)
}
```