package MaxDoubleSliceSum

import (
	"math"
)

func Solution(A []int) int {
	if len(A) < 4 {
		return 0
	}
	N := len(A) - 2
	forwardSum := make([]int, N)
	reverseSum := make([]int, N)

	//  0 ≤ X < Y < Z < N,
	//  A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].
	// 			A : [ 3,  2, 6, -1,  4,  5, -1, 2]
	// forwardSum : [ 0,  2, 8,  7, 11, 16]
	// reverseSum : [14,  8, 9,  5,  0,  0]
	for i := 0; i < N-1; i++ {
		forwardVal := A[i+1]
		reverseVal := A[N-i]

		forwardSum[i+1] = int(math.Max(0, float64(forwardVal)+float64(forwardSum[i])))
		reverseSum[N-i-2] = int(math.Max(0, float64(reverseVal)+float64(reverseSum[N-i-1])))
	}

	combineMax := math.MinInt64
	for i := 0; i < N; i++ {
		combineMax = int(math.Max(float64(combineMax), float64(forwardSum[i])+float64(reverseSum[i])))
	}

	return combineMax
}
