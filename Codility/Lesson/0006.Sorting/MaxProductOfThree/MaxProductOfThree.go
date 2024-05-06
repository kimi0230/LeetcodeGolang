package MaxProductOfThree

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int) int {
	sort.Ints(A)
	aLen := len(A)
	return max(A[0]*A[1]*A[aLen-1], A[aLen-1]*A[aLen-2]*A[aLen-3])
}
