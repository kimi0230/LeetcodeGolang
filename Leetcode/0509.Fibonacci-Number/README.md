# [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)

## 題目

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

```
F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
```
Given n, calculate F(n).


Example 1:
```
Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
```

Example 2:
```
Input: n = 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
```

Example 3:
```
Input: n = 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
```

Constraints:
* 0 <= n <= 30

## 題目大意
斐波那契數列, 通常用 F(n) 表示

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
給定 N，計算 F(N)。

提示：0 ≤ N ≤ 30
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377 ,610, 987……

## 解題思路
遇到遞迴最好畫出遞迴樹

```
            f(20)
           /     \ 
        f(19)   f(18)
       ...           ...
      /  \           /  \   
    f(1) f(2)       f(1) f(2) 
```

這一題解法很多，大的分類是四種，遞歸，記憶化搜索(dp)，矩陣快速冪，通項公式。其中記憶化搜索可以寫 3 種方法，自底向上的，自頂向下的，優化空間複雜度版的。通項公式方法實質是求 a^b 這個還可以用快速冪優化時間複雜度到 O(log n) 。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0500~0599/0509.Fibonacci-Number/
* https://leetcode.com/problems/fibonacci-number/
