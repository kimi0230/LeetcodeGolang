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

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0035.Search-Insert-Position/
* https://leetcode-cn.com/problems/search-insert-position/