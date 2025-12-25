package FrogRiverOne

func Solution(X int, A []int) int {
	intMap := make(map[int]bool)
	for i := 0; i < len(A); i++ {
		intMap[A[i]] = true
		if len(intMap) == X {
			return i
		}
	}
	return -1
}
