---
title:  1143.Longest Common Subsequence
tags:
  - Medium
  - Dynamic Programming
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [1143. Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/)

## 題目
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

Example 1:
```sh
Input: text1 = "abcde", text2 = "ace" 
Output: 3  
Explanation: The longest common subsequence is "ace" and its length is 3.
```

Example 2:

```sh
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
```

Example 3:

```sh
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
```

Constraints:

* 1 <= text1.length, text2.length <= 1000
* text1 and text2 consist of only lowercase English characters.

## 題目大意
給兩個string 求出, 這兩個string 的最長公共子序列的長度, 如果不存在返回0.
譬如 str1="abcde", str2="aceb", 輸出為3, 因為最長公共子序列是"ace"

## 解題思路
1. 暴力解法, 用遞迴 dp(i,j) 表示 s1[0..i]和s2[0..j]中最長公共子序列的長度, 
   如果s1[i]==s2[j], 說明這個公共字符一定在lcs中, 如果知道了s1[0..i-1]和s2[0..j-1]中的lcs長度, 再加1就是s1[0..i]和s2[0..j]中lcs的長度
   
    ```go
    if (str[i] == str2[j]) {
        dp(i,j) = dp(i-1,j-1)+1
    }
    ```

    如果s1[i]!=s2[j], 說明這兩個字符至少有一個不在lcs中,
    ```go
    if (str[i] != str2[j]){
        dp(i,j) = max( dp(i-1,j) , dp(i,j-1))
    }
    ```

    ```python
    def longestCommonSubsequence(str1,str2) ->int:
        def dp(i,j):
            # 空的base code
            if i == -1 or j == -1:
                return 0
            if str[i] == str2[j]:
                # 找到一個lcs中的元素
                return dp(i-1, j-1)+1
            if str[i] != str2[j]:
                # 至少有一個字符不在lcs中, 都試一下,看誰能讓lcs最長
                return max( dp(i-1,j) , dp(i,j-1))
        return dp(len(str1)-1,len(str2)-1)
    ```

2. DP優化

```c++
int longestCommonSubsequence(string str1, string str2) {
    int m = str1.size(), n = str2.size();
    // 定義對s1[0..i-1] 和 s2[0..j-1], 他們的lcs長度是dp[i][j]
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));
    // base case: dp[0][...] = dp[..][0] = 0, 已初始化

    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            // 狀態轉移邏輯
            if (str1[i - 1] == str2[j - 1]) {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                // 不用判斷 dp[i - 1][j-1] , 因為此永遠為三者中最小的, max根本不可能取到它
                dp[i][j] = max(dp[i][j - 1], dp[i - 1][j])
            }
        }
    }
    return dp[m][n];
}
```

## 來源
* https://leetcode.com/problems/longest-common-subsequence/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/1143.Longest-Common-Subsequence/main.go

```go
package longestcommonsubsequence

func LongestCommonSubsequence(text1 string, text2 string) int {
	var dp func(int, int) int
	dp = func(i, j int) int {
		if i == -1 || j == -1 {
			return 0
		}
		if text1[i] == text2[j] {
			return dp(i-1, j-1) + 1
		}
		if text1[i] != text2[j] {
			return max(dp(i-1, j), dp(i, j-1))
		}
		return 0
	}
	return dp(len(text1)-1, len(text2)-1)
}

func LongestCommonSubsequenceDP(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/1143.Longest-Common-Subsequence
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkLongestCommonSubsequence-8                  100         737158262 ns/op               0 B/op          0 allocs/op
BenchmarkLongestCommonSubsequenceDP-8            2355297               491.3 ns/op           912 B/op          8 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/1143.Longest-Common-Subsequence 75.400s
```