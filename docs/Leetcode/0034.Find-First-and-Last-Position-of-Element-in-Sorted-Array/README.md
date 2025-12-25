---
title: 0034. Find First and Last Position of Element in Sorted Array
subtitle: "https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/"
date: 2024-01-22T21:20:00+08:00
lastmod: 2024-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0034.Find-First-and-Last-Position-of-Element-in-Sorted-Array"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Medium
  - Find First and Last Position of Element in Sorted Array
  - Array
  - Binary Search
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
# [0034. Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/)

## 題目

## 題目大意


## 解題思路

## Big O

* 時間複雜 : `O(log(n))`
* 空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
* https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0034.Find-First-and-Last-Position-of-Element-in-Sorted-Array/main.go

```go
package findfirstandlastpositionofelementinsortedarray

// 時間複雜 O(log(n)), 空間複雜 O(1)
func searchRange(nums []int, target int) []int {
	result := []int{}

	// first index
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right) >> 1) // [0,1] => 0
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid
		}
	}
	if left == right && nums[left] == target {
		result = append(result, left)
	} else {
		result = append(result, -1)
	}

	// last index
	left, right = 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right+1) >> 1) // [0,1] => 1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if left == right && nums[left] == target {
		result = append(result, left)
	} else {
		result = append(result, -1)
	}
	return result
}

```

##  Benchmark

```sh

```