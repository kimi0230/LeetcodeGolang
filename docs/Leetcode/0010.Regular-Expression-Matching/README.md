---
title: 0010.Regular-Expression-Matching
subtitle: "https://leetcode.com/problems/regular-expression-matching/description/"
date: 2025-04-11T20:59:36+08:00
lastmod: 2025-04-11T20:59:36+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Regular-Expression-Matching"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Hard
  - String
  - Dynamic Programming
  - Recursion
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

# [0010.Regular-Expression-Matching](https://leetcode.com/problems/regular-expression-matching/description/)

## 題目
Given an input string s&nbsp;and a pattern p, implement regular expression matching with support for &#39;.&#39; and &#39;*&#39; where:


	&#39;.&#39; Matches any single character.​​​​
	&#39;*&#39; Matches zero or more of the preceding element.


The matching should cover the entire input string (not partial).

&nbsp;
Example 1:


Input: s = &quot;aa&quot;, p = &quot;a&quot;
Output: false
Explanation: &quot;a&quot; does not match the entire string &quot;aa&quot;.


Example 2:


Input: s = &quot;aa&quot;, p = &quot;a*&quot;
Output: true
Explanation: &#39;*&#39; means zero or more of the preceding element, &#39;a&#39;. Therefore, by repeating &#39;a&#39; once, it becomes &quot;aa&quot;.


Example 3:


Input: s = &quot;ab&quot;, p = &quot;.*&quot;
Output: true
Explanation: &quot;.*&quot; means &quot;zero or more (*) of any character (.)&quot;.


&nbsp;
Constraints:


	1 &lt;= s.length&nbsp;&lt;= 20
	1 &lt;= p.length&nbsp;&lt;= 20
	s contains only lowercase English letters.
	p contains only lowercase English letters, &#39;.&#39;, and&nbsp;&#39;*&#39;.
	It is guaranteed for each appearance of the character &#39;*&#39;, there will be a previous valid character to match.

## 題目大意


## 解題思路


## 來源
* https://leetcode.com/problems/regular-expression-matching/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0010.Regular-Expression-Matching/main.go



## Benchmark


