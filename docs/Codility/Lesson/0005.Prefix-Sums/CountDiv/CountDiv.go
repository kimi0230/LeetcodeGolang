package CountDiv

import "math"

// 時間: O(n)
func SolutionBurst(A int, B int, K int) int {
	result := 0
	for i := A; i <= B; i++ {
		if i%K == 0 {
			result++
		}
	}

	return result
}

// 時間:O(1) 空間: O(1)
func Solution(A int, B int, K int) int {
	result := 0
	if A%2 == 0 {
		result = 1
	}
	return int(math.Floor(float64(B/K))) - int(math.Floor(float64(A/K))) + result
}
