---
title: 300.Longest Increasing Subsequence
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
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/300.Longest-Increasing-Subsequence/main.go

```go