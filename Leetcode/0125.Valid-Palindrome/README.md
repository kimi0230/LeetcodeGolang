---
title: 0125. Valid Palindrome
subtitle: "https://leetcode.com/problems/valid-palindrome/"
date: 2023-10-13T16:11:00+08:00
lastmod: 2023-10-13T16:11:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Valid Palindrome"
license: ""
images: []

tags: [LeetCode, Go, Easy/Medium/Hard, Valid Palindrome]
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
# [0125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)

## 題目

## 題目大意


## 解題思路

## Big O
時間複雜 : `O(n)`
空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/valid-palindrome
* https://leetcode.cn/problems/valid-palindrome

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0125.Valid-Palindrome/main.go

```go
package validpalindrome

import (
	"strings"
	"unicode"
)

// 時間複雜 O(n), 空間複雜 O(1)
func IsPalindrome(s string) bool {
	// 將字符轉成小寫,
	s = strings.ToLower(s)

	// 使用雙指針, 一左一右相向而行, 判斷回文
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isChar(s[left]) {
			left++
		}
		for left < right && !isChar(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isChar(c byte) bool {
	if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) {
		return true
	}
	return false
}

func isCharAZ09(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
```

##  Benchmark

```sh

```