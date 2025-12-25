---
title: 0005. Longest Palindromic Substring
subtitle: "https://leetcode.com/problems/remove-nth-node-from-end-of-list/"
date: 2024-03-23T21:08:00+08:00
lastmod: 2024-03-23T21:08:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0005. Longest Palindromic Substring"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Medium
  - DP
  - Amazon
  - Microsoft
  - Google
  - Adobe
  - Facebook
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

# [0005.Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## 題目
Given a string s, return the longest palindromic substring in s.

A string is called a palindrome string if the reverse of that string is the same as the original string.


Example 1:
```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

Example 2:
```
Input: s = "cbbd"
Output: "bb"
```

Constraints:
* 1 <= s.length <= 1000
* s consist of only digits and English letters.


## 題目大意
給你一個字符串 s，找到 s 中最長的回文子串。

## 解題思路
* 每一個字符本身都是回文
* 長度為 2, 且首尾字符相同則為回文
* 長度>=3, 如果頭尾相同, 則去掉頭尾後可看是合是回文. 如果頭尾不同則不是回文

![](0005.LongestPalindromicString.jpg)

## 來源
* https://leetcode.com/problems/longest-palindromic-substring/
* https://leetcode.cn/problems/longest-palindromic-substring/solutions/1348874/s-by-xext-5zk3/

## 解答
```go
package longestpalindromicsubstring

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	result := s[0:1]

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 每個字符本身都是回文
	}
	for length := 2; length <= len(s); length++ {
		for start := 0; start < len(s)-length+1; start++ {
			end := start + length - 1
			if s[start] != s[end] {
				// 字頭字尾不同, 不是回文
				continue
			} else if length < 3 {
				// 長度為2且字頭字尾相同, 則為回文
				dp[start][end] = true
			} else {
				// 狀態轉移 : 去掉字頭字尾, 判斷是否還是回文
				dp[start][end] = dp[start+1][end-1]
			}
			if dp[start][end] && (end-start+1) > len(result) {
				result = s[start : end+1]
			}
		}
	}
	return result
}
```