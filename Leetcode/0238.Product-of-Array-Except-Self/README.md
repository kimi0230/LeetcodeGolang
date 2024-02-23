---
title: 238. Product of Array Except Self
subtitle: "https://leetcode.com/problems/product-of-array-except-self/description/"
date: 2024-02-19T18:01:00+08:00
lastmod: 2024-02-19T18:01:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0238.Product-of-Array-Except-Self"
license: ""
images: []

tags: [LeetCode, Go, Medium, Product of Array Except Self, Array, Prefix Sum, Blind75]
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
# [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/description/)

## 題目
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

 

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
 

Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
 

Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
## 題目大意


## 解題思路

## Big O
時間複雜 : `O(n)`
空間複雜 : ``

## 來源
* https://leetcode.com/problems/product-of-array-except-self/description/
* https://leetcode.cn/problems/product-of-array-except-self/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0238.Product-of-Array-Except-Self/main.go

```go
package productofarrayexceptself

// 時間複雜 O(), 空間複雜 O()
func productExceptSelf(nums []int) []int {
	result, left, right := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	// left 為左側所有的成績
	// 索引為'0' 的元素, 因為左側沒有元素,所以left[0]=1
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[i]
	}
	return result
}

func productExceptSelf2(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct = rightProduct * nums[i]
	}

	return result
}

```

##  Benchmark

```sh
go test -benchmem -run=none LeetcodeGolang/Leetcode/0238.Product-of-Array-Except-Self -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0238.Product-of-Array-Except-Self
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkProductExceptSelf-8            12772174               158.4 ns/op            96 B/op          3 allocs/op
BenchmarkProductExceptSelf2-8           32292304                63.74 ns/op           32 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0238.Product-of-Array-Except-Self       4.228s
```