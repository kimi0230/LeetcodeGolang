---
title: 0300.Longest Increasing Subsequence
tags: Medium, Dynamic Programming
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [300.Longest-Increasing-Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)

## 題目

Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

Example 1:
```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

Example 2:
```
Input: nums = [0,1,0,3,2,3]
Output: 4
```

Example 3:
```
Input: nums = [7,7,7,7,7,7,7]
Output: 1
 ```

Constraints:

* 1 <= nums.length <= 2500
* -104 <= nums[i] <= 104
 

Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?

## 題目大意
給你一個整事的數組 nums, 找到其中最長遞增子序列的長度
* 子序列不一定是連續的
* 子串一定是連續的

## 解題思路
### 方法一: DP
dp[i]表示nums[i]這個數結尾的最長遞增子序列的長度
譬如要找dp[5]的直, 也就是球nums[5]為結尾的最長遞增子序列
只要找到結尾比nums[5]小的子序列, 然後把 nums[5]接到最後,
就可以形成新的遞增子序列, 而這個新子序列長度+1

### 方法二: DP + 二分搜尋
patience sorting (耐心排序)
給一組 [6, 3 ,5 ,10, 11, 2, 9, 14, 13, 7, 4, 8, 12]
只能把數字小的放在數字大的堆上, 
如果當前數字太大沒有可以放置的堆, 則新建一個堆, 
如果當前排有很多堆可以選擇, 則選擇最左邊的那一堆放, 因為這樣可以確保堆頂是有序的(2,4,7,8,12)
6   5   10  11  14
3   4   9   8   13
2       7       12

堆數就是最長遞增子序列的長度
[3,5,7,8,12]

## 來源
* https://leetcode.com/problems/longest-increasing-subsequence/
  
## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0300.Longest-Increasing-Subsequence/main.go

```go
package longestincreasingsubsequence

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//  DP 解法 O(n^2)
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for idx := 0; idx < len(dp); idx++ {
		dp[idx] = 1
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 找到前面比現在結尾還小的子序列
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	// fmt.Println(dp)
	return res
}

func LengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	res := 0

	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				// 找到前面比現在結尾還小的子序列
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}
	// fmt.Println(dp)
	return res
}

//  DP + 二分搜尋:patience sorting
func LengthOfLISPatience(nums []int) int {
	top := make([]int, len(nums))

	// 牌堆數初始化為0
	piles := 0

	for i := 0; i < len(nums); i++ {
		poker := nums[i]

		// 搜尋左側邊界的二元搜尋
		left, right := 0, piles
		// fmt.Printf("i:%d\tL:%d\tR:%d\n", i, left, right)
		for left < right {
			mid := left + (right-left)/2
			if top[mid] > poker {
				// 現在的牌比堆小, 所小範圍
				right = mid
			} else if top[mid] < poker {
				// 現在的牌比堆大
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 沒有找到堆, 建立一個新的堆
		if left == piles {
			piles++
		}
		// 再把這張牌放在堆的頂端
		top[left] = poker
	}
	return piles
}
```

```
go test -benchmem -run=none LeetcodeGolang/Leetcode/300.Longest-Increasing-Subsequence -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/300.Longest-Increasing-Subsequence
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkLengthOfLIS-8                  18300901                64.77 ns/op           64 B/op          1 allocs/op
BenchmarkLengthOfLIS2-8                 17410562                68.98 ns/op           80 B/op          1 allocs/op
BenchmarkLengthOfLISPatience-8          20768851                58.47 ns/op           64 B/op          1 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/300.Longest-Increasing-Subsequence      3.812s
```