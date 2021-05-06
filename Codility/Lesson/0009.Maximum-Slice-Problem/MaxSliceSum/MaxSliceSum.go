package MaxSliceSum

import "math"

func Solution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	result := math.MinInt64
	sum := math.MinInt64
	for i := 0; i < len(A); i++ {
		sum = int(math.Max(float64(A[i]), float64(A[i])+float64(sum)))
		result = int(math.Max(float64(sum), float64(result)))
	}
	return result
}
