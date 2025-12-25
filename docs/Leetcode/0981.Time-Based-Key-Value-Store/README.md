---
title: 0981.Time Based Key-Value Store
subtitle: "https://leetcode.com/problems/time-based-key-value-store/description/"
date: 2024-01-22T21:20:00+08:00
lastmod: 2024-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0981.Time-Based-Key-Value-Store"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Medium
  - Time Based Key-Value Store
  - Hash Table
  - String
  - Binary Search
  - Design
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
# [0981.Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/description/)

## 題目

## 題目大意


## 解題思路

## Big O

* 時間複雜 : ``
* 空間複雜 : ``

## 來源
* https://leetcode.com/problems/time-based-key-value-store/description/
* https://leetcode.cn/problems/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0981.Time-Based-Key-Value-Store/main.go

```go
package timebasedkeyvaluestore

import "sort"

// 時間複雜 O(), 空間複雜 O()
type element struct {
	key       string
	value     string
	timestamp int
}

type TimeMap struct {
	elements map[string][]element
}

func Constructor() TimeMap {
	return TimeMap{elements: make(map[string][]element)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.elements[key] = append(this.elements[key], element{
		key:       key,
		value:     value,
		timestamp: timestamp,
	})
}

// O(log n)
func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.elements[key]; !ok {
		return ""
	}
	// 二分法找出符合條件的最左邊
	index := sort.Search(len(this.elements[key]), func(i int) bool {
		return this.elements[key][i].timestamp > timestamp
	})
	if index == 0 {
		return ""
	}

	index--
	return this.elements[key][index].value
}

```

##  Benchmark

```sh

```