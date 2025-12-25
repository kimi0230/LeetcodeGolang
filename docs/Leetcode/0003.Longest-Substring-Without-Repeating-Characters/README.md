---
title: 0003.Longest Substring Without Repeating Characters
subtitle: "https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/933491221/"
date: 2023-01-22T21:20:00+08:00
lastmod: 2023-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0003.Longest Substring Without Repeating Characters"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Medium
  - Longest Substring Without Repeating Characters
  - Sliding Window
  - Array
  - Hash Table
  - String
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
# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 题目
Given a string, find the length of the longest substring without repeating characters.

Example 1:

```c
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

Example 2:

```c
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```c
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## 題目大意

在一個字符串重尋找沒有重複字母的最長子串。

## 解題思路
這一題和第 438 題，第 3 題，第 76 題，第 567 題類似，用的思想都是"滑動窗口"。

滑動窗口的右邊界不斷的右移，只要沒有重複的字符，就持續向右擴大窗口邊界。一旦出現了重複字符，就需要縮小左邊界，直到重複的字符移出了左邊界，然後繼續移動滑動窗口的右邊界。以此類推，每次移動需要計算當前長度，並判斷是否需要更新最大長度，最終最大的值就是題目中的所求。

O(n)

**用空間換取時間, map紀錄已出現過的字符, 如果map效能不好時可使用數組(Slice)來改善**

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters/
* https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0003.Longest-Substring-Without-Repeating-Characters/Longest-Substring-Without-Repeating-Characters.go

```go
package longestSubstringwithoutrepeatingcharacters

// LengthOfLongestSubstring 暴力解
func LengthOfLongestSubstring(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	tmpLen := 1
	var maxLen = 1

	for i := 1; i < slength; i++ {
		// 往前找前幾個視窗
		j := i - tmpLen

		for ; j < i; j++ {
			if s[j] == s[i] { // 如果相同，那麼和S[J]到S[I-1]中間的肯定不相同，所以可以直接計算得到
				tmpLen = i - j
				break
			}
		}

		if j == i { // 都不相同
			tmpLen++
		}

		if tmpLen > maxLen {
			maxLen = tmpLen
		}
	}

	return maxLen
}

// LengthOfLongestSubstringMap 用map 紀錄是否重複.
func LengthOfLongestSubstringMap(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	charMap := make(map[byte]bool)
	maxLen, left, right := 0, 0, 0

	for left < slength {
		if ok := charMap[s[right]]; ok {
			// 有找到
			charMap[s[left]] = false
			left++
		} else {
			charMap[s[right]] = true
			right++
		}
		if maxLen < right-left {
			maxLen = right - left
		}
		if (left+maxLen) >= slength || right >= len(s) {
			break
		}
	}

	return maxLen
}

// LengthOfLongestSubstringBit 用map效能不好時可使用數組改善
func LengthOfLongestSubstringBit(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	// ASCII 0~255
	charMap := [256]bool{}
	maxLen, left, right := 0, 0, 0
	for left < slength {
		if ok := charMap[s[right]]; ok {
			// 有找到
			charMap[s[left]] = false
			left++
		} else {
			charMap[s[right]] = true
			right++
		}

		if maxLen < right-left {
			maxLen = right - left
		}
		if left+maxLen >= slength || right >= len(s) {
			break
		}
	}
	return maxLen
}

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0003.Longest-Substring-Without-Repeating-Characters
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkLengthOfLongestSubstring-8             66143602                19.08 ns/op            0 B/op          0 allocs/op
BenchmarkLengthOfLongestSubstringMap-8           2524627               397.8 ns/op             0 B/op          0 allocs/op
BenchmarkLengthOfLongestSubstringBit-8          65099846                21.37 ns/op            0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0003.Longest-Substring-Without-Repeating-Characters     4.193s
```