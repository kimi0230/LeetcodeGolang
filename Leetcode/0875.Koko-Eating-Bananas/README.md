---
title: 875. Koko Eating Bananas
subtitle: "https://leetcode.com/problems/koko-eating-bananas/description/"
date: 2024-03-28T18:03:00+08:00
lastmod: 2024-03-28T18:03:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0875.Koko-Eating-Bananas"
license: ""
images: []

tags: [LeetCode, Go, Medium, Koko Eating Bananas, Array, Binary Search]
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
# [875. Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/description/)

## 題目

## 題目大意
珂珂喜歡吃香蕉。 這裡有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。 警衛已經離開了，將在 H 小時後回來。
可以決定她吃香蕉的速度 K （單位：根/小時）。 每個小時，她將會選擇一堆香蕉，從中吃掉 K 根。 如果這堆香蕉少於 K 根，她將吃掉這堆的所有香蕉，然後這一小時內不會再吃更多的香蕉。
珂珂喜歡慢慢吃，但仍然想在警衛回來前吃掉所有的香蕉。
返回她可以在 H 小時內吃掉所有香蕉的最小速度 K（K 為整數）。

## 解題思路
這一題可以用二分搜索來解答。 在 [0 , max(piles)] 的範圍內搜索，二分的過程都是常規思路。 判斷是否左右邊界如果劃分的時候需要注意題目中給的限定條件。 當香蕉個數小於 k 的時候，那個小時不能再吃其他香蕉了。

## Big O

* 時間複雜 : `O（n log m）`，其中 n 是香蕉堆的數量，m 是香蕉堆中香蕉數量的最大值
* 空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/koko-eating-bananas/description/
* https://leetcode.cn/problems/koko-eating-bananas/description/
* https://books.halfrost.com/leetcode/ChapterFour/0800~0899/0875.Koko-Eating-Bananas/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0875.Koko-Eating-Bananas/main.go

```go
package kokoeatingbananas

// 時間複雜 O（n log m）, 空間複雜 O(1)
func minEatingSpeed(piles []int, h int) int {
	// 找出最大的香蕉堆
	left, right := 1, maxPile(piles)
	for left < right {
		// mid 代表消耗香蕉的速度(k)
		mid := int(uint(left+right) >> 1)
		if executeTime(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 假設消耗速度k, 算出要花多少時間
func executeTime(piles []int, k int) int {
	time := 0
	for _, pile := range piles {
		time += pile / k
		if pile%k > 0 {
			time++
		}
		// time += int(math.Ceil(float64(pile) / float64(k)))
	}
	return time
}

func maxPile(piles []int) int {
	result := 0
	for _, pile := range piles {
		if result < pile {
			result = pile
		}
	}
	return result
}

```

##  Benchmark

```sh

```