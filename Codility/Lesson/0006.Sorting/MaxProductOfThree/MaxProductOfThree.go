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
	alen := len(A)
	return max(A[0]*A[1]*A[alen-1], A[alen-1]*A[alen-2]*A[alen-3])
}
