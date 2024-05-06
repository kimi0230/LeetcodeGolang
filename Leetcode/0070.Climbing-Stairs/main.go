package climbingstairs

// 時間複雜 O(n), 空間複雜 O(n)
func ClimbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func ClimbStairsCache(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		if val := dp[i]; val == 0 {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}

func ClimbStairsRecursive(n int) int {
	dp := make([]int, n+1)
	// dp[0], dp[1] = 1, 1

	var climbClosure func(n int) int
	climbClosure = func(n int) int {
		if n <= 2 {
			return n
		}
		dp[n] = climbClosure(n-1) + climbClosure(n-2)
		return dp[n]
	}
	return climbClosure(n)
}
