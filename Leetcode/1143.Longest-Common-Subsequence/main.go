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

// DP table 優化
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
