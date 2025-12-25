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

tags:
  - LeetCode
  - Go
  - Easy
  - Valid Palindrome
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
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
 

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.

## 題目大意

## 解題思路

左右指針


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

// 時間複雜 O(), 空間複雜 O()
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

func IsPalindromeStrBuilder(s string) bool {
	s = strings.ToLower(s)
	strLen := len(s)
	// 將字符轉成小寫, 並濾掉空格和標點符號等字符
	var sb strings.Builder
	for i := 0; i < strLen; i++ {
		c := rune(s[i])
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			sb.WriteRune(c)
		}
	}

	sbStr := sb.String()

	// 使用雙指針, 一左一右相向而行, 判斷回文
	left, right := 0, len(sbStr)-1
	for left < right {
		if sbStr[left] != sbStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0125.Valid-Palindrome
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkIsPalindrome-8                  3840048               552.2 ns/op            32 B/op          1 allocs/op
BenchmarkIsPalindromeStrBuilder-8        3164848               439.0 ns/op            88 B/op          4 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0125.Valid-Palindrome   5.242s
```