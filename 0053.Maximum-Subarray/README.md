# [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

## 題目

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

**Example:**
    Input: [-2,1,-3,4,-1,2,1,-5,4],
    Output: 6
    Explanation: [4,-1,2,1] has the largest sum = 6.

**Follow up:**
If you have figured out the O(*n*) solution, try coding another solution using the divide and conquer approach, which is more subtle.

## 題目大意

給定一個整數數組 nums ，找到一個具有最大和的連續子數組（子數組最少包含一個元素），返回其最大和。

## 解題思路

- 這一題可以用 DP 求解也可以不用 DP。
- 題目要求輸出數組中某個區間內數字之和最大的那個值。 `dp[i]` 表示`[0,i]` 區間內各個子區間和的最大值，狀態轉移方程是`dp[i] = nums[i] + dp[i-1] (dp[i- 1] > 0)`，`dp[i] = nums[i] (dp[i-1] ≤ 0)`。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0053.Maximum-Subarray/
* https://leetcode-cn.com/problems/maximum-subarray/