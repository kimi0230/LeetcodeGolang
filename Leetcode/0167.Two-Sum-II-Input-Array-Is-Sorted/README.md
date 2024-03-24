---
title: 0167.Two Sum II Input Array Is Sorted
subtitle: "https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/"
date: 2023-10-21T13:54:00+08:00
lastmod: 2023-10-21T13:54:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0167.Two-Sum-II-Input-Array-Is-Sorted"
license: ""
images: []

tags: [LeetCode, Go, Medium, Two Sum II Input Array Is Sorted, Array, Two Pointers, Binary Search]
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
# [0167.Two Sum II Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

## 題目
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

 

Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
 

Constraints:

2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers is sorted in non-decreasing order.
-1000 <= target <= 1000
The tests are generated such that there is exactly one solution.

## 題目大意
找出兩個數之和等於 target 的兩個數位，要求輸出它們的下標。 注意一個數位不能使用 2 次。 下標從小到大輸出。 假定題目一定有一個解。

## 解題思路

## Big O
1. 使用1. Two Sun來解
   1. O(n^2) 的時間複雜度和 O(1)的空間複雜度暴力
   2. 藉助哈希表使用 O(n) 的時間複雜度和 O(n) 的空間複雜度求解

2. 雙指針
時間複雜度： O(n)，其中 n 是陣列的長度。 兩個指標移動的總次數最多為 n 次。
空間複雜度： O(1)

時間複雜 : ``
空間複雜 : ``

## 來源
* https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
* https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0167.Two-Sum-II-Input-Array-Is-Sorted/main.go

```go
package twosumiiinputarrayissorted

// 雙指針. 時間複雜度： O(n)，其中 n 是陣列的長度。 兩個指標移動的總次數最多為 n 次。
// 空間複雜度： O(1)

func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

// 使用 O(n^2) 的時間複雜度和 O(1)的空間複雜度暴力
// 或者藉助哈希表使用 O(n) 的時間複雜度和 O(n) 的空間複雜度求解

func TwoSum2(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))

	for i, v := range numbers {
		if j, ok := m[target-v]; ok {
			return []int{j + 1, i + 1}
		}
		m[v] = i
	}
	return nil
}


```

##  Benchmark

```sh
go test -benchmem -run=none LeetcodeGolang/Leetcode/0167.Two-Sum-II-Input-Array-Is-Sorted -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0167.Two-Sum-II-Input-Array-Is-Sorted
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkTwoSum-4       35161128                39.74 ns/op           16 B/op          1 allocs/op
BenchmarkTwoSum2-4      19309680                70.21 ns/op           16 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0167.Two-Sum-II-Input-Array-Is-Sorted   2.866s
```