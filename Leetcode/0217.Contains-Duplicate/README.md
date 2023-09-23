---
title: 0217.Contains-Duplicate
subtitle: "https://leetcode.com/problems/contains-duplicate/description/"
date: 2023-01-22T21:20:00+08:00
lastmod: 2023-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Contains Duplicate"
license: ""
images: []

tags: [LeetCode, Go, Easy, Contains Duplicate]
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
# [0217.Contains-Duplicate](https://leetcode.com/problems/contains-duplicate/description/)

## 題目
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.


Example 1:
```
Input: nums = [1,2,3,1]
Output: true
```

Example 2:
```
Input: nums = [1,2,3,4]
Output: false
```

Example 3:
```
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

```

Constraints:

* 1 <= nums.length <= 105
* -109 <= nums[i] <= 109

## 題目大意


## 解題思路


## 來源
* https://leetcode.com/problems/contains-duplicate/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0217.Contains-Duplicate/main.go

```go
package containsduplicate

func ContainsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		if numsMap[v] {
			return true
		}
		numsMap[v] = true
	}
	return false
}

```

##  Benchmark

```sh

```