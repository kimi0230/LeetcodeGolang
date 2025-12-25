package coinchange

import (
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var memo = map[int]int{}

func dp(coins []int, n int) int {
	// 查詢備忘錄 避免重複
	if _, vok := memo[n]; vok {
		return memo[n]
	}

	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	res := math.MaxInt
	for _, coin := range coins {
		subproblem := dp(coins, n-coin)
		if subproblem == -1 {
			continue
		}
		res = min(res, 1+subproblem)
	}
	if res != math.MaxInt {
		memo[n] = res
	} else {
		memo[n] = -1
	}
	return memo[n]
}

// 備忘錄 遞迴 時間複雜O(kn).
func CoinChangeMemoryTableRecursion(coins []int, amount int) int {
	memo = map[int]int{}
	return dp(coins, amount)
}

// dp array 迭代解法
func CoinChangeDP(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	// 初始化, 湊成 amount 金額的硬幣 最多就 amount 個(全都用1元), 所以 amount+1相當於正的無窮
	// dp[]的定義: 當目標金額為i時, 至少需要dp[i]枚硬幣湊出
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	// 外層 for 循環遍歷所有狀態的所有取值
	for i := 1; i <= amount; i++ {
		// 內層 for 循環遍歷求所有選擇的最小值
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			// 狀態轉移
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func CoinChange(coins []int, n int) int {
	var dpClosure func(n int) int
	dpClosure = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := math.MaxInt
		for _, coin := range coins {
			subproblem := dpClosure(n - coin)
			if subproblem == -1 {
				continue
			}
			res = min(res, 1+subproblem)
		}
		if res != math.MaxInt {
			return res
		} else {
			return -1
		}
	}
	return dpClosure(n)
}
