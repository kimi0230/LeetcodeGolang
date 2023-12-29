package mincostclimbingstairs

// 時間複雜 O(n), 空間複雜 O(1)
func MinCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[len(cost)-2], dp[len(cost)-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 使用了兩個變數cur和last，
// cur 變數存儲從第 i-2 步到達第 i 步的最小花費。
// last 變數存儲從第 i-1 步到達第 i 步的最小花費。
func MinCostClimbingStairsOptimize(cost []int) int {
	var cur, last int
	for i := 2; i < len(cost)+1; i++ {
		if last+cost[i-1] > cur+cost[i-2] {
			cur, last = last, cur+cost[i-2]
		} else {
			cur, last = last, last+cost[i-1]
		}
	}
	return last
}
