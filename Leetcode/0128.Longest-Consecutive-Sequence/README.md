---
title: 128. Longest Consecutive Sequence
subtitle: "https://leetcode.com/problems/longest-consecutive-sequence/description/"
date: 2024-02-20T20:53:00+08:00
lastmod: 2024-02-20T20:53:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0128.Longest-Consecutive-Sequence"
license: ""
images: []

tags: [LeetCode, Go, Medium, Longest Consecutive Sequence, Array, Hash Table, Union Find]
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
# [128. Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/description/)

## 題目

## 題目大意


## 解題思路

## Big O
時間複雜 : `O(n)`
空間複雜 : `O(n)`

## 來源
* https://leetcode.com/problems/longest-consecutive-sequence/description/
* https://leetcode.cn/problems/longest-consecutive-sequence/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0128.Longest-Consecutive-Sequence/main.go

```go
package longestconsecutivesequence

// 時間複雜 O(), 空間複雜 O()
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}
	result := 0
	for v := range m {
		// 如果沒找到該數字的前一個數字, 則把該數字刀做連續序列的第一個數
		if _, ok := m[v-1]; !ok {
			length := 1
			for _, exit := m[v+length]; exit; _, exit = m[v+length] {
				length++
			}
			if result < length {
				result = length
			}
		}
	}
	return result
}

```

##  Benchmark

```sh

```