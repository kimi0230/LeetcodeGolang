---
title: 0322.Coin Change
tags: Medium, Dynamic Programming
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [322. Coin Change](https://leetcode.com/problems/coin-change/) 
## 題目
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

 

Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0
 

Constraints:

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104


## 題目大意
給妳 k 種面值的硬幣, 面值分別為 c1,c2 ...,ck, 每種硬幣的數量無限.
再給你一個總金額. 求出**最少** 需要幾枚硬幣湊出這個金額, 如果不能湊出 return -1.

## 解題思路
`dp(n)` 的定義: 輸入一個目標金額n, 返回湊出目標金額n的最少硬幣數量

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0322.Coin-Change/
* https://leetcode.com/problems/coin-change/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0322.Coin-Change/Coin-Change.go

```go
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
		// 內層 for 循環遍歷球所有選擇的最小值
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
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

```

```shell
go test -benchmem -run=none LeetcodeGolang/Leetcode/0322.Coin-Change -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0322.Coin-Change
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkCoinChange-4                             273376              4452 ns/op               0 B/op          0 allocs/op
BenchmarkCoinChangeDP-4                         11071686               128.1 ns/op            96 B/op          1 allocs/op
BenchmarkCoinChangeMemoryTableRecursion-4       57663068                23.69 ns/op            0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0322.Coin-Change        4.194s
```


```
amount = 11. k=[1,2,5]

amount 0 1 2 3 4 5 ... 9 10 11
index  0 1 2 3 4 5 ... 9 10 11
dp     0 1 1 2 2 1 ... 3  2  3


dp[5] = 1+min(dp[5-1],dp[5-2],dp[5-5])
dp[11] = 1+min(dp[10],dp[9],dp[6])
```