package longestpalindromicsubsequence

/*
*
dp:方程表示從i到j的長度中，回文子串的長度
Base case:
a ->dp[i][i]=1
case 1:s[i]==s[j]

	a*****a ->dp[i][j]=dp[i+1][j-1]+2

case 2:s[i]!=s[j]

	ab****b dp[i][j]=dp[i+1][i]
	a****ab dp[i][j]=dp[i][j-1]
*/
func LongestPalindromeSubseq(s string) int {
	n := len(s)
	// 初始化  dp table : [][]int{}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		// Base case
		dp[i][i] = 1
	}

	// 反向遍歷確保正確的狀態轉移
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
			// fmt.Printf("[%d][%d] = %d \n", i, j, dp[i][j])
		}
	}
	return dp[0][n-1]
}

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func max[T numbers](a T, b T) T {
	if a > b {
		return a
	}
	return b
}
