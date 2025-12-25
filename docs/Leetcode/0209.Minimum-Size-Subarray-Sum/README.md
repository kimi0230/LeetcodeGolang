---
title: 0209. Minimum Size Subarray Sum
tags:
  - Medium
  - Sliding Window
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)

## 题目

Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example 1:

```c
Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
```

```
Input: s = 4, nums = [1,4,4]
Output: 1
```

```
Input: s = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
```

Follow up:
  
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n). 

## 題目大意

給定一個整型數組和一個數字 s，找到數組中最短的一個連續子數組，使得連續子數組的數字之和 sum>=s，返回最短的連續子數組的返回值。

## 解題思路

這一題的解題思路是用滑動窗口。在滑動窗口[i,j]之間不斷往後移動，如果總和小於s，就擴大右邊界j，不斷加入右邊的值，直到sum > s，之和再縮小i 的左邊界，不斷縮小直到sum < s，這時候右邊界又可以往右移動。以此類推。

## 進階

如果你已經實現 O(n) 時間複雜度的解法, 請嘗試設計一個 O(n log(n)) 時間複雜度的解法。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0209.Minimum-Size-Subarray-Sum/
* https://leetcode-cn.com/problems/minimum-size-subarray-sum/
* https://mp.weixin.qq.com/s/UrZynlqi4QpyLlLhBPglyg

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0209.Minimum-Size-Subarray-Sum/Minimum-Size-Subarray-Sum.go

```go
package minimumsizesubarraysum

import (
	"math"
	"sort"
)

// MinSubArrayLenBurst : 暴力解 時間複雜 O(n^2), 空間複雜 O(1)
func MinSubArrayLenBurst(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}

	minSize := math.MaxInt32
	for i := 0; i < lens; i++ {
		sum := 0
		for j := i; j < lens; j++ {
			sum += nums[j]
			if sum >= target {
				minSize = min(minSize, j-i+1)
			}
		}
	}
	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

// MinSubArrayLenSlidingWindow : 滑動視窗 時間複雜 O(n), 空間複雜 O(1)
// 滑動窗口的精妙之處在於根據當前子序列和大小的情況，不斷調節子序列的起始位置。從而將O(n^2)的暴力解法降為O(n)
func MinSubArrayLenSlidingWindow(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}
	minSize := math.MaxInt32

	start, end, sum := 0, 0, 0
	for end < lens {
		sum += nums[end]
		for sum >= target {
			minSize = min(minSize, end-start+1)
			sum -= nums[start]
			start++
		}
		end++

	}

	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

/*
MinSubArrayLenBinarysearch : 前缀和 + 二分查找  O(nlog(n))
為了使用二分查找，需要額外創建一個數組 sums 用於存儲數組nums 的前綴和，其中 sums[i] 表示從 nums[0] 到 nums[i−1] 的元素和。
得到前綴和之後，對於每個開始下標i，可通過二分查找得到大於或等於 i 的最小下標 bound，
使得 sums[bound]-sums[i-1] >= s，
並更新子數組的最小長度（此時子數組的長度是 bound -(i-1) )。
C++ 的 lower_bound，Java 的 Arrays.binarySearch

因為這道題保證了數組中每個元素都為正，所以前綴和一定是遞增的，這一點保證了二分的正確性。如果題目沒有說明數組中每個元素都為正，這裡就不能使用二分來查找這個位置了。
*/
func MinSubArrayLenBinarysearch(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}
	minSize := math.MaxInt32

	// sums[i] 表示從 nums[0] 到 nums[i−1]
	sums := make([]int, lens+1)
	// 為了方便計算 size = lens + 1
	// sums[0] = 0, 前0個的元素和
	// sums[1] = nums[0] , 前1個元素和為 nums[0]
	for i := 1; i <= lens; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	// fmt.Println("sums", sums)
	for i := 1; i <= lens; i++ {
		/*
			target      7
			input       2  3  1  2  4  3
			-----------------------------------
			i           0  1  2  3  4  5  6
			-----------------------------------
			sums        0  2  5  6  8 12 15	// input 的前 n 個元素和
			tmpTarge       7  9 12 13 15 19	// tmpTarge = target + sums[i-1]
			bound          4  5  5  6  6  7	// bound = sort.SearchInts(sums, tmpTarge)
			bound-(i-1)    4  4  3  3  2  2
			minSize        4  4  3  3  2  2
		*/
		tmpTarge := target + sums[i-1]
		bound := sort.SearchInts(sums, tmpTarge)
		// if bound < 0 {
		// 	bound = -bound - 1
		// }
		if bound <= lens {
			minSize = min(minSize, bound-(i-1))
		}
		// fmt.Println("i:", i, " tmpTarge", tmpTarge, "  bound:", bound, "  bound-(i-1)", bound-(i-1), " minSize:", minSize)
	}

	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```