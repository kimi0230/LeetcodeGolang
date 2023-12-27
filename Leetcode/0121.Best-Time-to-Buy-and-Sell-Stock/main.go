package besttimetobuyandsellstock

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func max[T numbers](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// 時間複雜 O(n), 空間複雜 O(1)
// Slide windows
func MaxProfit(prices []int) int {
	left, right := 0, 1
	maxProfit := 0
	for right < len(prices) {
		if prices[left] < right {
			profit := prices[right] - prices[left]
			maxProfit = max(maxProfit, profit)
		} else {
			left = right
		}
		right++
	}
	return maxProfit
}

// 時間複雜 O(n), 空間複雜 O(1)
// DP : 最佳解
func MaxProfitDP(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0] // 初始最低價格為第一天的價格
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// 如果當前價格比最低價格低，更新最低價格
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// 否則計算當前價格賣出時的利潤，並更新最大利潤
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

// 時間複雜 O(), 空間複雜 O()
// DP
func MaxProfitDP2(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		if i-1 == -1 {
			// base case
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}
