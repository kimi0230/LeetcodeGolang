---
title: 412. Fizz Buzz
subtitle: "https://leetcode.com/problems/fizz-buzz/description/"
date: 2024-02-14T19:37:00+08:00
lastmod: 2024-02-14T19:37:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0412.Fizz-Buzz"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Easy
  - Fizz Buzz
  - Facebook
  - Microsoft
  - Apple
  - string
  - math
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
# [412. Fizz Buzz](https://leetcode.com/problems/fizz-buzz/description/)

## 題目
`Facebook`, `Microsoft`, `Apple`
Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
 

Example 1:

Input: n = 3
Output: ["1","2","Fizz"]
Example 2:

Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]
Example 3:

Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
 

Constraints:

1 <= n <= 104

## 題目大意


## 解題思路

## Big O
時間複雜 : `O(n)`
空間複雜 : `O(n)`

## 來源
* https://leetcode.com/problems/fizz-buzz/description/
* https://leetcode.cn/problems/fizz-buzz/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0412.Fizz-Buzz/main.go

```go
package fizzbuzz

import "strconv"

// 時間複雜 O(), 空間複雜 O()
func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func fizzBuzz2(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		var str string
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if len(str) <= 0 {
			str = strconv.Itoa(i)
		}
		result = append(result, str)
	}
	return result
}

// 最佳
func fizzBuzz3(n int) []string {
	var result []string
	fizz := 0
	buzz := 0
	for i := 1; i <= n; i++ {
		fizz++
		buzz++
		if fizz == 3 && buzz == 5 {
			result = append(result, "FizzBuzz")
			fizz = 0
			buzz = 0
		} else if fizz == 3 {
			result = append(result, "Fizz")
			fizz = 0
		} else if buzz == 5 {
			result = append(result, "Buzz")
			buzz = 0
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

```

##  Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0412.Fizz-Buzz
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkFizzBuzz-4      5918809               287.1 ns/op           112 B/op          3 allocs/op
BenchmarkFizzBuzz2-4     5024536               223.8 ns/op           112 B/op          3 allocs/op
BenchmarkFizzBuzz3-4     5406643               196.3 ns/op           112 B/op          3 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0412.Fizz-Buzz  5.507s
```