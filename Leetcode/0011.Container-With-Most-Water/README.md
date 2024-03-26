---
title: 0011.Container With Most Water
subtitle: "https://leetcode.com/problems/container-with-most-water/description/"
date: 2024-03-24T16:31:00+08:00
lastmod: 2024-03-24T16:31:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0011.Container-With-Most-Water"
license: ""
images: []

tags: [LeetCode, Go, Medium, Container With Most Water, Amazon, Microsoft, Adobe, Facebook, Google, Array, Two Pointers, Greedy]
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
# [0011.Container With Most Water](https://leetcode.com/problems/container-with-most-water/description/)

## 題目

## 題目大意


## 解題思路
利用雙指針, 找出最小的高並乘上兩指針距離, 得出面積
當左指針高度比右指針高度高時, 將右指針往左移, 反之將左指針往右移.

## Big O

* 時間複雜 : `O(n)`
* 空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/container-with-most-water/description/
* https://leetcode.cn/problems/container-with-most-water/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0011.Container-With-Most-Water/main.go

```go
package containerwithmostwater

// 時間複雜 O(n), 空間複雜 O(1)
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		tmp := (right - left) * min(height[left], height[right])
		result = max(result, tmp)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```

##  Benchmark

```sh

```