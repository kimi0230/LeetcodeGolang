package TapeEquilibrium

import "math"

func Solution(A []int) int {

	totalSum := 0
	for _, v := range A {
		totalSum += v
	}

	leftSum := A[0]
	rightSum := totalSum - leftSum
	result := math.MaxInt32
	for i := 1; i < len(A); i++ {
		tmpDiff := int(math.Abs(float64(rightSum) - float64(leftSum)))
		if tmpDiff < result {
			result = tmpDiff
		}
		rightSum -= A[i]
		leftSum += A[i]
	}

	if result == math.MaxInt32 {
		result = 0
	}

	return result
}
