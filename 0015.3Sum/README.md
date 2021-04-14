# [15. 3Sum](https://leetcode.com/problems/3sum/)

## 題目

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

```c
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

## 題目大意

給定一個數組，要求在這個數組中找出 3 個數之和為 0 的所有組合。

## 解題思路

用 map 提前計算好任意 2 個數字之和，保存起來，可以將時間複雜度降到 O(n^2)。這一題比較麻煩的一點在於，最後輸出解的時候，要求輸出不重複的解。數組中同一個數字可能出現多次，同一個數字也可能使用多次，但是最後輸出解的時候，不能重複。例如[-1，-1，2] 和[2, -1, -1]、[-1, 2, -1] 這3 個解是重複的，即使-1 可能出現100 次，每次使用的-1 的數組下標都是不同的。

這裡就需要去重和排序了。 map 記錄每個數字出現的次數，然後對 map 的 key 數組進行排序，最後在這個排序以後的數組裡面掃，找到另外 2 個數字能和自己組成 0 的組合。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0015.3Sum/
* https://leetcode-cn.com/problems/3sum/