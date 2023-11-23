---
title: 0242.Valid-Anagram
subtitle: "https://leetcode.com/problems/valid-anagram/description/"
date: 2023-09-23T21:20:00+08:00
lastmod: 2023-09-23T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Valid Anagram"
license: ""
images: []

tags: [LeetCode, Go, Easy, Valid Anagram]
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
# [0242.Valid-Anagram](https://leetcode.com/problems/valid-anagram/description/)
驗證回文串

## 題目
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
 

Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
 

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
> 我們使用 rune 數據類型來處理 Unicode 字元，並使用 map[rune]int 來統計這些字符的出現次數。其餘的邏輯保持不變。

## 題目大意

## 解題思路
只要先把所有字符轉換成小寫，並過濾掉空格和標點這類字符，然後對剩下的字符執行 [數組雙指針技巧匯總](https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-48c1d/shuang-zhi-fa4bd/) 中提到的兩端向中心的雙指針演算法即可

## 來源
* https://leetcode.com/problems/valid-anagram/description/
* https://leetcode.cn/problems/valid-anagram/description/
* [數組雙指針技巧匯總](https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-48c1d/shuang-zhi-fa4bd/) 

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0242.Valid-Anagram/main.go

```go
package validanagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := make(map[rune]int, len(s))

	for _, v := range s {
		record[v]++
	}
	for _, v := range t {
		record[v]--
		if record[v] < 0 {
			return false
		}
	}
	return true
}

func IsAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := make(map[rune]int, len(s))

	for _, v := range s {
		record[v]++
	}
	for _, v := range t {
		record[v]--
	}

	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0242.Valid-Anagram
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkIsAnagram-8             6703497               276.3 ns/op             0 B/op          0 allocs/op
BenchmarkIsAnagram2-8            3660909               318.9 ns/op            48 B/op          2 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0242.Valid-Anagram      4.498s
```