package MaxProfit

import (
	"math"
)

func Solution(A []int) int {
	minBuyPrice := math.MaxFloat64
	maxProfit := 0.0

	for _, v := range A {
		minBuyPrice = math.Min(minBuyPrice, float64(v))
		maxProfit = math.Max(maxProfit, float64(v)-minBuyPrice)
	}

	return int(maxProfit)
}
