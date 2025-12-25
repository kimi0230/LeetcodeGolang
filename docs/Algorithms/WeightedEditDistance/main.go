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
