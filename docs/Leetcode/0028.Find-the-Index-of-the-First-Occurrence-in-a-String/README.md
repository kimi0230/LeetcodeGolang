---
title: 0028. Find the Index of the First Occurrence in a String
subtitle: "https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/"
date: 2024-03-18T23:20:00+08:00
lastmod: 2024-03-18T23:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0028.Find-the-Index-of-the-First-Occurrence-in-a-String"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Easy
  - Find the Index of the First Occurrence in a String
  - Two Pointers
  - String
  - String Matching
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
# [0028. Find the Index of the First Occurrence in a String](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/)

## 題目

## 題目大意


## 解題思路

## Big O

* 時間複雜 : ``
* 空間複雜 : ``

## 來源
* https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
* https://leetcode.cn/problems/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0028.Find-the-Index-of-the-First-Occurrence-in-a-String/main.go

```go
package findtheindexofthefirstoccurrenceinastring

// 暴力解
// 時間複雜 O(M*N), 空間複雜 O()
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	index := 0
	for i := 0; i <= (haystackLen - needleLen); i++ {
		j := 0
		for j = 0; j < needleLen; j++ {
			if haystack[i+j] == needle[j] {
				index = i
			} else {
				break
			}
		}
		if j == needleLen {
			return index
		}
	}
	return -1
}

// Slice 解法
func strStrSlice(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if haystackLen == 0 || haystackLen < needleLen {
		return -1
	}
	if needleLen == 0 {
		return 0
	}
	for i := 0; i <= (haystackLen - needleLen); i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}

```

##  Benchmark

```sh

```