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

// 時間複雜 O(n), 空間複雜 O(1)
// 優化
// 使用了兩個變數cur和last，
// cur 變數存儲從第 i-2 步到達第 i 步的最小花費。
// last 變數存儲從第 i-1 步到達第 i 步的最小花費。
// 比較兩種選擇的花費：
// 從第2階開始（i := 2），一直迭代到最後一階（i < len(cost)+1）。
// 在每一步驟中：
// 比較兩種選擇的花費：
//
//	選擇從上一階跳到當前階的花費：last + cost[i-1]
//	選擇從上上階跳到當前階的花費：cur + cost[i-2]
//
// 選擇花費較小的方案：
//
//	如果 last + cost[i-1] 更小，則將 last 的值更新為 cur + cost[i-2]，並將 cur 的值更新為 last（即上一步的最小花費）。
//	否則，將 last 的值更新為 last + cost[i-1]，並將 cur 的值更新為 last（即上一步的最小花費）。
func MinCostClimbingStairsOptimize(cost []int) int {
	var cur, last int
	for i := 2; i < len(cost)+1; i++ {
		// 選擇從上一階跳到當前階的花費：last + cost[i-1]
		// 選擇從上上階跳到當前階的花費：cur + cost[i-2]
		if last+cost[i-1] > cur+cost[i-2] {
			cur, last = last, cur+cost[i-2]
		} else {
			cur, last = last, last+cost[i-1]
		}
		// fmt.Printf("%-d | %-d | %-d\n", i, cur, last)
	}
	return last
}
