---
title: 3223.Minimum-Length-of-String-After-Operations
subtitle: "https://leetcode.com/problems/minimum-length-of-string-after-operations/description/"
date: 2025-04-11T20:56:29+08:00
lastmod: 2025-04-11T20:56:29+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Minimum-Length-of-String-After-Operations"
license: ""
images: []

tags: [LeetCode, Go, Medium, Hash Table, String, Counting]
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
mapbox:
share:
  enable: true
comment:
  enable: true
library:
  css:
  js:
seo:
  images: []
---

# [3223.Minimum-Length-of-String-After-Operations](https://leetcode.com/problems/minimum-length-of-string-after-operations/description/)

## 題目
You are given a string s.

You can perform the following process on s any number of times:


	Choose an index i in the string such that there is at least one character to the left of index i that is equal to s[i], and at least one character to the right that is also equal to s[i].
	Delete the closest occurrence of s[i] located to the left of i.
	Delete the closest occurrence of s[i] located to the right of i.


Return the minimum length of the final string s that you can achieve.

&nbsp;
Example 1:


Input: s = &quot;abaacbcbb&quot;

Output: 5

Explanation:
We do the following operations:


	Choose index 2, then remove the characters at indices 0 and 3. The resulting string is s = &quot;bacbcbb&quot;.
	Choose index 3, then remove the characters at indices 0 and 5. The resulting string is s = &quot;acbcb&quot;.



Example 2:


Input: s = &quot;aa&quot;

Output: 2

Explanation:
We cannot perform any operations, so we return the length of the original string.


&nbsp;
Constraints:


	1 &lt;= s.length &lt;= 2 * 105
	s consists only of lowercase English letters.

## 題目大意


## 解題思路


## 來源
* https://leetcode.com/problems/minimum-length-of-string-after-operations/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/3223.Minimum-Length-of-String-After-Operations/main.go



## Benchmark


