# [46. Permutations](https://leetcode.com/problems/permutations/)

## 題目
Given a collection of **distinct** integers, return all possible permutations(排列).
**Example:**
    Input: [1,2,3]
    Output:
    [
      [1,2,3],
      [1,3,2],
      [2,1,3],
      [2,3,1],
      [3,1,2],
      [3,2,1]
    ]

## 題目大意
給定一個沒有重複數字的序列，返回其所有可能的全排列。

## 解題思路
- 求出一個數組的排列組合中的所有排列，用 DFS 深搜即可。
這個問題可以看作有 ñ 個排列成一行的空格，我們需要從左往右依此填入題目給定的 ñ個數，每個數只能使用一次。
那麼很直接的可以想到一種窮舉的算法，即從左往右每一個位置都依此嘗試填入一個數，
看能不能填完這ñ 個空格，在程序中我們可以用「回溯法」來模擬這個過程
回溯法：
一種通過探索所有可能的候選解來找出所有的解的算法。如果候選解被確認不是一個解（或者至少不是最後一個解），
回溯算法會通過在上一步進行一些變化拋棄該解，即回溯並且再次嘗試。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/permutations/solution/quan-pai-lie-by-leetcode-solution-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0046.Permutations/
* https://leetcode-cn.com/problems/permutations/