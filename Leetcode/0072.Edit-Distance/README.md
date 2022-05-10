---
title: 0072. Edit Distance
tags: Hard, Dynamic Programming
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [0072. Edit Distance](https://leetcode.com/problems/edit-distance/)

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

```
dp = [
	[0,1,2,3,4,5],
	[1,1,2,2,3,4],
	[2,2,1,2,3,4],
	[3,3,2,2,2,3]
]
```

| word1 \ word2 |  "" |  h  |  o  |  r  |  s  |  e  |
|---------------|-----|-----|-----|-----|-----|-----|
| ""            |  0  |  1  |  2  |  3  |  4  |  5  |
| r             |  1  |  1  |  2  |  2  |  3  |  4  |
| o             |  2  |  2  |  1  |  2  |  3  |  4  |
| s             |  3  |  3  |  2  |  2  |  2  |  3  |

dp(i,j) 返回值, 就是 word1[0..i] 和 word2[0..j]的最小編輯距離
dp(1,0) "ro" , "h" 最小編輯距離 2

| dp[i-1][j-1] | dp[i-1][j] |
|--------------|------------|
| dp[i][j-1]   |  dp[i][j]  |

| 替換/跳過 | 刪除 |
|--------------|------------|
| 插入  |   |

## 來源
* https://leetcode.com/problems/edit-distance/
* https://labuladong.gitee.io/algo/3/25/77/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0072.Edit-Distance/main.go

```go
package editdistance

import "fmt"

// 遞迴 (暴力解)
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
			// word1[0..i] 和 word2[0..j]的最小編輯距離等於 word1[0..i-1] 和 word2[0..j-1]
			// 本來就相等所以不需要任何操作
			// 也就是說 dp(i,j)等於 dp(i-1,j-1)
			return dp(i-1, j-1)
		} else {
			return min(
				dp(i, j-1)+1,   // insert: 直接在 word1[i]中插入一個和word2[j]一樣的字符, 那麼word2[j]就被匹配了,往前j, 繼續和i對比, 操作次數+1
				dp(i-1, j)+1,   // delete: 直接把 word1[i] 這個字符串刪除, 往前 i 繼續和 j 對比, 操作次數+1
				dp(i-1, j-1)+1, // replace: 直接把 word1[i] 替換成 word2[j], 這樣他們就匹配了, 同時往前 i, j 繼續對比, 操作次數+1
			)
		}
	}

	return dp(len(word1)-1, len(word2)-1)
}

// Memo優化
func MinDistanceMemo(word1 string, word2 string) int {
	var dp func(int, int) int
	memo := map[string]int{}
	dp = func(i, j int) int {
		key := fmt.Sprintf("%d,%d", i, j)
		// 查詢備忘錄 避免重複
		if _, ok := memo[key]; ok == true {
			return memo[key]
		}

		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if word1[i] == word2[j] {
			// word1[0..i] 和 word2[0..j]的最小編輯距離等於 word1[0..i-1] 和 word2[0..j-1]
			// 本來就相等所以不需要任何操作
			// 也就是說 dp(i,j)等於 dp(i-1,j-1)
			memo[key] = dp(i-1, j-1)
		} else {
			memo[key] = min(
				dp(i, j-1)+1,   // insert: 直接在 word1[i]中插入一個和word2[j]一樣的字符, 那麼word2[j]就被匹配了,往前j, 繼續和i對比, 操作次數+1
				dp(i-1, j)+1,   // delete: 直接把 word1[i] 這個字符串刪除, 往前 i 繼續和 j 對比, 操作次數+1
				dp(i-1, j-1)+1, // replace: 直接把 word1[i] 替換成 word2[j], 這樣他們就匹配了, 同時往前 i, j 繼續對比, 操作次數+1
			)
		}
		return memo[key]
	}

	return dp(len(word1)-1, len(word2)-1)
}

// DP table 優化, DP table 是自底向上求解, 遞迴是自頂向下求解
func MinDistanceDP(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 初始化  dp table : [][]int{}
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// 向上求解
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i][j-1]+1,   // insert
					dp[i-1][j]+1,   // delete
					dp[i-1][j-1]+1, // replace
				)
			}
		}
	}
	return dp[m][n]
}

type Number interface {
	int | int64 | float64
}

func min[T Number](vars ...T) T {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}


```


```sh
go test -benchmem -run=none LeetcodeGolang/Leetcode/0072.Edit-Distance -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0072.Edit-Distance
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMinDistance-8            398260              3748 ns/op               0 B/op          0 allocs/op
BenchmarkMinDistanceMemo-8        102272             10796 ns/op            2211 B/op         69 allocs/op
BenchmarkMinDistanceDP-8         1944886               794.2 ns/op           688 B/op          9 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0072.Edit-Distance      5.717s
```