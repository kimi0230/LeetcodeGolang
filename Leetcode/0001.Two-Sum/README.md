# [1. Two Sum](https://leetcode.com/problems/two-sum/)
## 题目
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

## 題目大意

在數組中找到 2 個數之和等於給定值的數字，結果返回 2 個數字在數組中的下標。

## 解題思路

這道題最優的做法時間複雜度是 O(n)。

順序掃描數組，對每一個元素，在 map 中找能組合給定值的另一半數字，如果找到了，直接返回 2 個數字的下標即可。如果找不到，就把這個數字存入 map 中，等待掃到“另一半”數字的時候，再取出來返回結果。

O(n)

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0001.Two-Sum/
* https://leetcode-cn.com/problems/two-sum/