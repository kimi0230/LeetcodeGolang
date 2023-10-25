---
title: 0020. Valid Parentheses
subtitle: "https://leetcode.com/problems/valid-parentheses/"
date: 2023-01-22T21:20:00+08:00
lastmod: 2023-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0020.Valid-Parentheses"
license: ""
images: []

tags: [LeetCode, Go, Easy, Valid Parentheses]
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
# [0020. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

## 題目

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

## 題目大意


## 解題思路

## Big O
時間複雜 : ``
空間複雜 : ``

## 來源
* https://leetcode.com/problems/valid-parentheses/
* https://leetcode.cn/problems/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0020.Valid-Parentheses/main.go

```go
package validparentheses

type Stack struct {
	runes []rune
}

func NewStack() *Stack {
	return &Stack{runes: []rune{}}
}

func (s *Stack) Push(str rune) {
	s.runes = append(s.runes, str)
}

func (s *Stack) Pop() rune {
	str := s.runes[len(s.runes)-1]
	s.runes = s.runes[:len(s.runes)-1]
	return str
}

// 時間複雜 O(n), 空間複雜 O(n)
func IsValid(s string) bool {
	runeStack := NewStack()
	for _, v := range s {
		// fmt.Println(string(v))
		if v == '(' || v == '[' || v == '{' {
			runeStack.Push(v)
		} else if (v == ')' && len(runeStack.runes) > 0 && runeStack.runes[len(runeStack.runes)-1] == '(') || (v == ']' && len(runeStack.runes) > 0 && runeStack.runes[len(runeStack.runes)-1] == '[') || (v == '}' && len(runeStack.runes) > 0 && runeStack.runes[len(runeStack.runes)-1] == '{') {
			runeStack.Pop()
		} else {
			return false
		}
	}
	return len(runeStack.runes) == 0
}
```

##  Benchmark

```sh

```