---
title: 74. Search a 2D Matrix
subtitle: "https://leetcode.com/problems/search-a-2d-matrix/description/"
date: 2024-03-24T22:10:00+08:00
lastmod: 2024-03-24T22:10:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0074.Search-a-2D-Matrix"
license: ""
images: []

tags: [LeetCode, Go, Medium, Search a 2D Matrix, Array, Binary Search, Matrix, Facebook, Amazon, Microsoft, Bloomberg]
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
# [74. Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/description/)

## 題目

## 題目大意
編寫一個高效的演算法來判斷 m x n 矩陣中，是否存在一個目標值。 該矩陣具有如下特性：
* 每行中的整數從左到右按升序排列。
* 每行的第一个整数大于前一行的最后一个整数。

## 解題思路
可以將二維轉成一維矩陣

## Big O

* 時間複雜 : `O(log⁡mn)` 其中 m 和 n 分別是矩陣的行數和列數
* 空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/search-a-2d-matrix/description/
* https://leetcode.cn/problems/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0074.Search-a-2D-Matrix/main.go

```go
package searcha2dmatrix

// 時間複雜 O(logmn), 空間複雜 O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m <= 0 {
		return false
	}
	n := len(matrix[0])

	left, right := 0, m*n-1
	for left <= right {
		mid := int(uint(right+left) >> 1) //left + (right-left)/2
		val := getMatrix(matrix, mid)
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 將 mxn的二維陣列 映射成一維陣列
func getMatrix(matrix [][]int, index int) int {
	n := len(matrix[0])
	i, j := index/n, index%n
	return matrix[i][j]
}

```

##  Benchmark

```sh

```