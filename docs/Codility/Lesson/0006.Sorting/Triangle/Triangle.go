package Triangle

import "sort"

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sort.Ints(A)
	for i := 0; i < len(A)-2; i++ {
		if A[i+2] < A[i+1]+A[i] {
			return 1
		}
	}
	return 0
}
