# [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/description/)

## 題目

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.

Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
Constraints:

-10^9 <= nums1[i], nums2[i] <= 10^9
nums1.length == m + n
nums2.length == n

## 題目大意

合併兩個已經有序的數組，結果放在第一個數組中，第一個數組假設空間足夠大。要求算法時間複雜度足夠低。

## 解題思路

為了不大量移動元素，就要從2個數組長度之和的最後一個位置開始，依次選取兩個數組中大的數，從第一個數組的尾巴開始往頭放，只要循環一次以後，就生成了合併以後的數組了。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0088.Merge-Sorted-Array/
* https://leetcode-cn.com/problems/merge-sorted-array/