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