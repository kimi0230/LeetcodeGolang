---
title: 0078. Subsets
subtitle: "https://leetcode.com/problems/subsets/description/"
date: 2023-11-21T17:30:00+08:00
lastmod: 2023-11-21T17:30:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0078.Subsets"
license: ""
images: []

tags: [LeetCode, Go, Medium, Subsets]
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
# [0078. Subsets](https://leetcode.com/problems/subsets/description/)

## 題目
Given an integer array nums of unique elements, return all possible 
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

 

Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]
 

Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.

## 題目大意
給定一組不含重複元素的整數陣列 nums，返回該陣列所有可能的子集（冪集）。 說明：解集不能包含重複的子集。

## 解題思路
找出一個集合中的所有子集，空集也算是子集。 且陣列中的數位不會出現重複。 用 DFS 暴力枚舉即可。
這一題和第 90 題，第 491 題類似，可以一起解答和複習

## Big O
時間複雜 : `O(n^2)`
空間複雜 : `O(n)`

## 來源
* https://leetcode.com/problems/subsets/description/
* https://leetcode.cn/problems/subsets/description/
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0078.Subsets/
* https://github.com/suntong/lang/blob/master/lang/Go/src/ds/Recursion/LC-0B-Subsets-2.go

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0078.Subsets/main.go

```go
package subsets

// 時間複雜 O(n^2)	, 空間複雜 O(n)
func Subsets(nums []int) [][]int {
	path, result := []int{}, [][]int{}
	for i := 0; i <= len(nums); i++ {
		genSubsets(nums, i, 0, path, &result)
	}
	return result
}

func genSubsets(nums []int, elemSize, start int, path []int, result *[][]int) {
	if len(path) == elemSize {
		b := make([]int, len(path))
		copy(b, path)
		*result = append(*result, b)
		return
	}

	for k := start; k < len(nums); k++ {
		path = append(path, nums[k])
		genSubsets(nums, elemSize, k+1, path, result)
		path = path[:len(path)-1]
	}
}

// 時間複雜 O(n^2)	, 空間複雜 O(n)
func SubsetsDFS(nums []int) [][]int {
	path, result := []int{}, [][]int{}
	genSubsetsDFS(nums, path, &result)
	return result
}

func genSubsetsDFS(nums []int, path []int, result *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	*result = append(*result, b)
	for i := 0; i < len(nums); i++ {
		// n個 element 在slice中
		genSubsetsDFS(nums[i+1:], append(path, nums[i]), result)
	}
}

```

##  Benchmark

```sh
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=. -cover
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMaxEnvelopes-8          7799131               160.3 ns/op            88 B/op          3 allocs/op
BenchmarkMaxEnvelopes2-8         5800399               195.6 ns/op            80 B/op          4 allocs/op
PASS
coverage: 96.0% of statements
ok      LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes     3.726s
```