package Distinct

import "sort"

func Solution(A []int) int {
	sort.Ints(A)
	if len(A) == 0 {
		return 0
	}
	result := 1
	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1] {
			result++
		}
	}
	return result
}

func SolutionSet(A []int) int {
	set := make(map[int]struct{})
	var void struct{}

	for i := 0; i < len(A); i++ {
		if _, ok := set[A[i]]; !ok {
			set[A[i]] = void
		}
	}
	return len(set)
}
