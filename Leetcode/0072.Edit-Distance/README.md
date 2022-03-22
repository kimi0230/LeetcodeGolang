---
title: 0072. Edit Distance
tags: Hard, Dynamic Programming
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [75. Sort Colors](https://leetcode.com/problems/edit-distance/)

## 题目
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

* Insert a character
* Delete a character
* Replace a character
 

Example 1:

```
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
```

Example 2:

```
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation: 
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
```

Constraints:

* 0 <= word1.length, word2.length <= 500
* word1 and word2 consist of lowercase English letters.

## 題目大意
可以對一個字符串進行三種操作, 插入, 刪除, 替換
現在給你兩個字符串word1,word2, 計算出word1轉換成word2**最少**需要多少次操作

## 解題思路
https://mp.weixin.qq.com/s/ShoZRjM8OyvDbZwoXh6ygg
解決兩個字符串的動態規劃問題, 一班都是用兩個指針 i, j 分別指向兩個字符串的最後, 然後一步步往前走, 縮小問題的規模

```python
if word1[i] == word2[j]:
    skip
    i,j同時往前
else:
    # insert
    # delete
    # replace
```

## 來源
* https://leetcode.com/problems/edit-distance/


## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0072.Edit-Distance/main.go

```go
package editdistance

// 遞迴
func MinDistance(word1 string, word2 string) int {
	var dp func(int, int) int
	dp = func(i, j int) int {
		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		} else {
			return min(
				dp(i, j-1)+1,   // insert
				dp(i-1, j)+1,   // delete
				dp(i-1, j-1)+1, // replace
			)
		}
	}

	return dp(len(word1)-1, len(word2)-1)
}

type Number interface {
	int | int64 | float64
}

func min[nums Number](vars ...nums) nums {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

```