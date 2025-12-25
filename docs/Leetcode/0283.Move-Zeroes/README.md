---
title: 0283. Move Zeroes
subtitle: "https://leetcode.com/problems/move-zeroes/"
date: 2023-03-13T22:13:00+08:00
lastmod: 2023-03-13T22:13:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0283. Move Zeroes"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Easy
  - Move Zeroes
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
# [283. Move Zeroes](LEETCODELINK)

## 題目
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

 

Example 1:

```
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
```

Example 2:
```
Input: nums = [0]
Output: [0]
```

Constraints:
* 1 <= nums.length <= 104
* -231 <= nums[i] <= 231 - 1
 

Follow up: Could you minimize the total number of operations done?

## 題目大意
題目要求不能利用額外的輔助空間，將數字組中 0 元素都移到數字組的末尾，並與維持所有非 0 元素的相對位置。

## 解題思路
1. 非零交換數據, 左右指針都往右移
2. 零, 右指針右移

## 來源
* https://leetcode.com/problems/move-zeroes/
* https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0283.Move-Zeroes/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0283.Move-Zeroes/main.go

```go
package movezeroes

// 時間O(n^2)
func MoveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

// 時間O(n), 空間O(1)
func MoveZeroes2point(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0283.Move-Zeroes
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkMoveZeroes-4           145544966                9.164 ns/op           0 B/op          0 allocs/op
BenchmarkMoveZeroes2point-4     183269922                6.149 ns/op           0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0283.Move-Zeroes        3.975s
```