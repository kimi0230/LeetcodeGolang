---
title: Find Target Last Index
subtitle: ""
date: 2023-11-16T17:09:00+08:00
lastmod: 2023-11-16T17:09:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "在有序的array中 找出target在array中最後的index是什麼"
license: ""
images: []

tags:
  - Interview
  - Algorithms
  - Go
  - Easy
  - Find Target Last Index
  - Right Bound
  - Left Bound
categories: [Algorithms]

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

# Find Target Last Index

在有序的array中 找出target在array中最後的index是什麼
用二分搜尋法去找, RightBound

可參考
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0704.Binary-Search/main.go
---

## 方法一

```go
func Solution(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := int(uint(right+left) >> 1)
		if nums[mid] == target {
			// 繼續找, 縮小右邊
			if target == nums[right] {
				result = right
				break
			} else {
				right--
			}
		} else if nums[mid] < target {
			// 往右找
			left = mid + 1
		} else if nums[mid] > target {
			// 往左找
			right = mid - 1
		}
	}
	return result
}
```

--- 

## 方法二 遞迴

```go
func SolutionRecursive(nums []int, target int) int {
	left, right := 0, len(nums)-1
	return findTarget(nums, left, right, target)
}

func findTarget(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := int(uint(left+right) >> 1)
	if nums[mid] == target {
		if nums[right] == target {
			return right
		} else {
			return findTarget(nums, mid, right-1, target)
		}
	} else if nums[mid] < target {
		return findTarget(nums, mid+1, right, target)
	} else {
		return findTarget(nums, left, mid-1, target)
	}

}

```

``` sh
go test -benchmem -run=none LeetcodeGolang/Algorithms/Find_Target_Last_Index -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Algorithms/Find_Target_Last_Index
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkSolution-8             54216429                24.75 ns/op            0 B/op          0 allocs/op
BenchmarkSolutionRecursive-8    40756744                27.75 ns/op            0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Algorithms/Find_Target_Last_Index        2.537s
```

---

## RightBound

```go
// 有點類似 nums 大於 target的元素有幾個
func RightBound(nums []int, target int) (index int) {
	lenght := len(nums)
	if lenght <= 0 {
		return -1
	}
	left, right := 0, lenght-1

	for left <= right {
		// 除以2
		// mid := left + (right-left)>>1
		mid := int(uint(right+left) >> 1)
		if nums[mid] == target {
			// 注意:要繼續找右邊, 所以把左邊變大=mid+1
			left = mid + 1
		} else if nums[mid] < target {
			// 找右邊
			left = mid + 1
		} else if nums[mid] > target {
			// 找左邊
			right = mid - 1
		}
	}
	// 都沒找到 注意:right越界情況
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
```


## LeftBound

``` go
// 有點類似 nums 小於 target的元素有幾個
func LeftBound(nums []int, target int) (index int) {
	lenght := len(nums)
	if lenght <= 0 {
		return -1
	}
	left, right := 0, lenght-1

	for left <= right {
		// 除以2
		// mid := left + (right-left)>>1
		mid := int(uint(right+left) >> 1)
		if nums[mid] == target {
			// 要繼續找左邊, 所以把右邊變小
			right = mid - 1
		} else if nums[mid] < target {
			// 找右邊
			left = mid + 1
		} else if nums[mid] > target {
			// 找左邊
			right = mid - 1
		}
	}
	// 都沒找到 注意: left越界情況
	if left >= lenght || nums[left] != target {
		return -1
	}
	return left
}
```