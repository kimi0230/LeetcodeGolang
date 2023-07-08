---
title: "WeightedEditDistance"
subtitle: "WeightedEditDistance"
date: 2023-07-08T15:23:25+08:00
lastmod: 2023-07-08T15:23:25+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: ""
license: ""
images: []

tags: [Golang, WeightedEditDistance, Dynamic Programming]
categories: [Algorithms]

featuredImage: ""
featuredImagePreview: ""

hiddenFromHomePage: false
hiddenFromSearch: false
twemoji: false
lightgallery: true
ruby: true
fraction: true
fontawesome: true
linkToMarkdown: true
rssFullText: false

toc:
  enable: true
  auto: true
code:
  copy: true
  maxShownLines: 50
math:
  enable: false
  # ...
mapbox:
  # ...
share:
  enable: true
  # ...
comment:
  enable: true
  # ...
library:
  css:
    # someCSS = "some.css"
    # located in "assets/"
    # Or
    # someCSS = "https://cdn.example.com/some.css"
  js:
    # someJS = "some.js"
    # located in "assets/"
    # Or
    # someJS = "https://cdn.example.com/some.js"
seo:
  images: []
  # ...
---
# WeightedEditDistance

WeightedEditDistance 是一個計算帶有權重的編輯距離的函式。編輯距離是衡量兩個字串之間的相似度的指標，表示將一個字串轉換為另一個字串所需的最小操作數量。

在標準的編輯距離算法中，操作包括插入、刪除和替換字符。每個操作都被認為具有相同的代價。然而，在 `WeightedEditDistance` 中，每個字符的操作代價可以不同，並由一個權重映射表指定。

函式 `WeightedEditDistance` 接受兩個字符串 `word1` 和 `word2`，以及一個權重映射表 `weights`。該映射表將每個字符映射到其相應的權重值，用於計算操作的代價。

該函式使用動態規劃的方法計算編輯距離。它創建一個二維矩陣 `dp`，其中 `dp[i][j]` 表示將 `word1[:i]` 轉換為 `word2[:j]` 的最小操作代價。

算法的核心是遍歷 `dp` 矩陣並計算每個單元格的值。如果 `word1[i-1]` 等於 `word2[j-1]`，則表示兩個字符相等，不需要進行操作，所以 `dp[i][j]` 等於 `dp[i-1][j-1]`。否則，需要考慮插入、刪除和替換操作的代價，並取其中最小的作為 `dp[i][j]` 的值。

最終，函式返回 `dp[m][n]`，其中 `m` 和 `n` 分別為 `word1` 和 `word2` 的長度，表示將整個字串 `word1` 轉換為 `word2` 的最小操作代價。

使用 `WeightedEditDistance` 函式，您可以根據字符的權重值計算帶有自定義操作代價的編輯距離，以更好地反映兩個字串之間的相似性。

## 題目大意
weightededitdistance 雖然不是一個特定的LeetCode問題，但它涉及到一個概念：加權編輯距離（Weighted Edit Distance）。

加權編輯距離是指在兩個字串之間進行編輯操作（插入、刪除、替換）時，每個操作具有不同的成本或權重。該問題要求計算從一個字串轉換到另一個字串的最小總成本或權重。

## 解題思路
解決加權編輯距離問題的常用方法是使用動態規劃（Dynamic Programming）。

1. 創建一個二維數組dp，其中dp[i][j]表示將字串1的前i個字符轉換為字串2的前j個字符的最小加權編輯距離。

2. 初始化dp矩陣的第一行和第一列，分別表示將空字串轉換為字串1和字串2的成本，根據具體問題設置初始值。

3. 遍歷dp矩陣，計算每個dp[i][j]的值，根據以下三種情況進行選擇：

   - 如果字串1的第i個字符等於字串2的第j個字符，則dp[i][j]等於dp[i-1][j-1]，即不需要進行編輯操作，繼承前一個狀態的編輯距離。

   - 否則，dp[i][j]等於插入操作的成本加上dp[i][j-1]，刪除操作的成本加上dp[i-1][j]，替換操作的成本加上dp[i-1][j-1]，取這三種操作的最小值。

4. 最終，dp[m][n]（其中m和n分別為兩個字串的長度）即為兩個字串的最小加權編輯距離。


| 替換 /跳過 <br> dp[i-1][j-1] | 刪除 <br> dp[i-1][j] |
|------------------------------|----------------------|
| 插入 <br> dp[i][j-1]         | dp[i][j]             |

時間複雜度: 動態規劃的遍歷過程需要計算和填充dp矩陣的每個元素，因此時間複雜度為O(m*n)，其中m和n分別為兩個字串的長度。

空間複雜度: 需要使用一個二維數組dp來保存中間結果，因此空間複雜度為O(m*n)。

## 解答

```go
package weightededitdistance

import "fmt"

func WeightedEditDistance(word1, word2 string, weights map[rune]int) int {
	m, n := len(word1), len(word2)

	// 創建二維矩陣用於保存編輯距離
	// dp，其中 dp[i][j] 表示將 word1[:i] 轉換為 word2[:j] 的最小操作代價
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化第一列, base case
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + weights[rune(word1[i-1])]
	}

	// 初始化第一行
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + weights[rune(word2[j-1])]

	}
	// fmt.Println(dp)
	// 填充編輯距離矩陣
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 計算插入、刪除和替換操作的代價
				// insert: 直接在 word1[i]中插入一個和word2[j]一樣的字符, 那麼word2[j]就被匹配了,往前j, 繼續和i對比, 操作次數+1
				insertCost := dp[i][j-1] + weights[rune(word2[j-1])]
				deleteCost := dp[i-1][j] + weights[rune(word1[i-1])]
				replaceCost := dp[i-1][j-1] + weights[rune(word1[i-1])] + weights[rune(word2[j-1])]
				// 取最小的代價作為當前操作的編輯距離
				dp[i][j] = min(insertCost, deleteCost, replaceCost)
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

// 輔助函式，返回三個數字中的最小值
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

```

相關 [0072.Edit-Distance](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0072.Edit-Distance)
