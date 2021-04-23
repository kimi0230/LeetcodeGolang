package MinAvgTwoSlice

import "math"

func Solution(A []int) int {

	min := math.MaxFloat64
	minIndex := -1
	for i := 0; i < len(A)-1; i++ {
		// 2個數平均
		if i+1 < len(A) {
			tmp := (float64(A[i]) + float64(A[i+1])) / float64(2)
			if tmp < min {
				min = tmp
				minIndex = i
			}
		}
		// 3個數平均
		if i+2 < len(A) {
			tmp := (float64(A[i]) + float64(A[i+1]) + float64(A[i+2])) / float64(3)
			if tmp < min {
				min = tmp
				minIndex = i
			}
		}
	}

	return minIndex
}
