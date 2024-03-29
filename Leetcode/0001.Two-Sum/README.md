---
title: 0001.Two Sum
tags: Easy
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [1. Two Sum](https://leetcode.com/problems/two-sum/)
## 题目
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

## 題目大意

在數組中找到 2 個數之和等於給定值的數字, 結果返回 2 個數字在數組中的下標. 

## 解題思路
* 暴力解: 時間複雜 O(n^2), 空間複雜 O(1)

* 這道題最優的做法時間複雜度是 O(n)
順序掃描數組, 對每一個元素，在 map 中找能組合給定值的另一半數字, 如果找到了, 直接返回 2 個數字的下標即可. 如果找不到, 就把這個數字存入 map 中, 等待掃到“另一半”數字的時候, 再取出來返回結果. 

* 如果nums是有序 可以使用左右指針


## 來源
* https://leetcode-cn.com/problems/two-sum/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0001.Two-Sum/twosum.go

```go
package twosum

// 時間複雜 O(n^2), 空間複雜 O(1)
func Twosum(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func Twosum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return []int{0, 0}
}

// 如果nums是有序 可以使用左右指針
// func Twosum3(nums []int, target int) []int {
// 	left, right := 0, len(nums)-1
// 	sort.Ints(nums)

// 	for left < right {
// 		sum := nums[left] + nums[right]
// 		if sum == target {
// 			return []int{left, right}
// 		} else if sum < target {
// 			left++
// 		} else if sum > target {
// 			right--
// 		}
// 	}
// 	return []int{0, 0}
// }
```