---
title: 0035.Search Insert Position
tags:
  - Easy
  - Array
author: Kimi Tsai <kimi0230@gmail.com>
description:
---

# [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)
## 題目

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

**Example 1:**

    Input: [1,3,5,6], 5
    Output: 2

**Example 2:**

    Input: [1,3,5,6], 2
    Output: 1

**Example 3:**

    Input: [1,3,5,6], 7
    Output: 4

**Example 4:**

    Input: [1,3,5,6], 0
    Output: 0

## 題目大意

給定一個排序數組和一個目標值，在數組中找到目標值，並返回其索引。如果目標值不存在於數組中，返回它將會被按順序插入的位置。

你可以假設數組中無重複元素。

## 解題思路
- 給出一個已經從小到大排序後的數組，要求在數組中找到插入 target 元素的位置。
- 這一題是經典的二分搜索的變種題，在有序數組中找到最後一個比 target 小的元素。
- 只要看到面試題裡給出的數組是有序數組，都可以想一想是否可以使用二分法

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0035.Search-Insert-Position/
* https://leetcode-cn.com/problems/search-insert-position/
* [數組：每次遇到二分法，都是一看就會，一寫就廢](https://mp.weixin.qq.com/s/fCf5QbPDtE6SSlZ1yh_q8Q)

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0035.Search-Insert-Position/Search-Insert-Position.go

```go
package searchinsertposition

// 暴力解 時間複雜  O(n) 空間複雜 O(1)
func SearchInsertBurst(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

//二分法 時間複雜 O(log n) 空間複雜 O(1)
func SearchInsertBisection(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// / 防止溢出 同(left + right)/2
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	// 分別處理如下四種情況
	// targe在所有元素之前 [0, -1]
	// targe等於數組中某一個元素 return middle;
	// targe在數組中的位置 [left, right]，return right + 1
	// targe在數組所有元素之後的情況 [left, right]， return right + 1
	return right + 1
}

//二分法 時間複雜 O(log n) 空間複雜 O(1)
func SearchInsertBisection2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// / 防止溢出 同(left + right)/2
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
```