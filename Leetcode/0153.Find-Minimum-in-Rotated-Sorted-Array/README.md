---
title: 0153.Find-Minimum-in-Rotated-Sorted-Array
subtitle: "https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/"
date: 2024-03-29T00:27:00+08:00
lastmod: 2024-03-29T00:27:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0153.Find-Minimum-in-Rotated-Sorted-Array"
license: ""
images: []

tags: [LeetCode, Go, Medium, LEETCODETITLE,Amazon, Facebook, Microsoft, Adobe, Goldman Sachs]
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
# [0153.Find-Minimum-in-Rotated-Sorted-Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/)

## 題目

## 題目大意


## 解題思路

## Big O

* 時間複雜 : `O(logN)`
* 空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
* https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0153.Find-Minimum-in-Rotated-Sorted-Array/main.go

```go
package findminimuminrotatedsortedarray

// 時間複雜 O(logN), 空間複雜 O(1)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := int(uint(left+right) >> 1)
		if nums[mid] >= nums[left] {
			// mid 一定不是最小值了
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

```

##  Benchmark

```sh

```