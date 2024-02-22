---
title: 0015. 3Sum
subtitle: "https://leetcode.com/problems/3sum/description/"
date: 2024-02-22T23:04:00+08:00
lastmod: 2024-02-22T23:04:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0015.3Sum"
license: ""
images: []

tags: [LeetCode, Go, Medium, 3Sum, Array, Two Pointers, Sorting]
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
# [15. 3Sum](https://leetcode.com/problems/3sum/)

## 題目

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

```c
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

## 題目大意

給定一個數組，要求在這個數組中找出 3 個數之和為 0 的所有組合。

## 解題思路

用 map 提前計算好任意 2 個數字之和，保存起來，可以將時間複雜度降到 O(n^2)。這一題比較麻煩的一點在於，最後輸出解的時候，要求輸出不重複的解。數組中同一個數字可能出現多次，同一個數字也可能使用多次，但是最後輸出解的時候，不能重複。例如[-1，-1，2] 和[2, -1, -1]、[-1, 2, -1] 這3 個解是重複的，即使-1 可能出現100 次，每次使用的-1 的數組下標都是不同的。

這裡就需要去重和排序了。 map 記錄每個數字出現的次數，然後對 map 的 key 數組進行排序，最後在這個排序以後的數組裡面掃，找到另外 2 個數字能和自己組成 0 的組合。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0015.3Sum/
* https://leetcode-cn.com/problems/3sum/


## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0015.3Sum/3Sum.go

```go
package threesum

import (
	"sort"
)

// ThreeSumBurst : 暴力解 : O(n^3)
func ThreeSumBurst(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums) // O(n log n)
	for i := 0; i < len(nums); i++ {
		// 需要跟上一次不同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			// 需要跟上一次不同
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}

/*
ThreeSumDoublePoint : 最佳解, 排序 + 雙指針法 (滑動視窗) O(n^2)
1. 特判，對於數組長度 n，如果數組為 null 或者數組長度小於 3，返回 []。
2. 對數組進行排序。
3. 遍歷排序後數組：
  - 對於重複元素：跳過，避免出現重複解
  - 令左指針 L=i+1，右指針 R=n−1，當 L<R 時，執行循環：
  - 當nums[i]+nums[L]+nums[R]==0，執行循環，判斷左界和右界是否和下一位置重複，
    去除重複解。並同時將 L,R 移到下一位置，尋找新的解
  - 若和大於 0，說明 nums[R] 太大，R 左移
  - 若和小於 0，說明 nnums[L] 太小，L 右移
*/
func ThreeSumDoublePoint(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums) // O(n log n)
	start, end, addNum := 0, 0, 0
	for i := 1; i < len(nums)-1; i++ {
		start, end = 0, len(nums)-1
		if i > 1 && nums[i] == nums[i-1] {
			// 去掉重複
			start = i - 1
		}
		for start < i && end > i {
			if start > 0 && nums[start] == nums[start-1] {
				// 去掉重複
				start++
				continue
			}
			if end < (len(nums)-1) && nums[end] == nums[end+1] {
				// 去掉重複
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[i]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[i], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func ThreeSumHashTable(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums) // O(n log n)

	for i := 0; i < len(nums)-2; i++ {
		// 避免重複的起始元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		seen := make(map[int]bool)
		target := -nums[i] // 目標值為當前元素的相反數
		for j := i + 1; j < len(nums); j++ {
			complement := target - nums[j] // 找到與當前元素配對的目標元素
			if seen[complement] {
				result = append(result, []int{nums[i], complement, nums[j]})
				// 避免重複的配對元素
				for j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}
			}
			seen[nums[j]] = true
		}
	}
	return result
}

func ThreeSumTwoPointer(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, l, r := -nums[i], i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > target {
				r--
			} else if sum < target {
				l++
			}
		}
	}
	return result
}

```


```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0015.3Sum
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkThreeSumBurst-4                 9838000               121.4 ns/op            48 B/op          2 allocs/op
BenchmarkThreeSumDoublePoint-4           9069201               112.8 ns/op            48 B/op          2 allocs/op
BenchmarkThreeSumHashTable-4             7935907               147.1 ns/op            48 B/op          2 allocs/op
BenchmarkThreeSumTwoPointer-4           10888315               103.5 ns/op            48 B/op          2 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0015.3Sum       5.055s
```