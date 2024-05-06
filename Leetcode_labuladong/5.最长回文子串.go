// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	result := s[0:1]

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 每個字符本身都是回文
	}
	for length := 2; length <= len(s); length++ {
		for start := 0; start < len(s)-length+1; start++ {
			end := start + length - 1
			if s[start] != s[end] {
				// 字頭字尾不同, 不是回文
				continue
			} else if length < 3 {
				// 長度為2且字頭字尾相同, 則為回文
				dp[start][end] = true
			} else {
				// 狀態轉移 : 去掉字頭字尾, 判斷是否還是回文
				dp[start][end] = dp[start+1][end-1]
			}
			if dp[start][end] && (end-start+1) > len(result) {
				result = s[start : end+1]
			}
		}
	}
	return result
}

// @lc code=end

