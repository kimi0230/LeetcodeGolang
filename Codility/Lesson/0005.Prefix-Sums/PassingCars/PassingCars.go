package PassingCars

func Solution(A []int) int {
	addBase, result := 0, 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			addBase++
		} else {
			result += addBase
		}
	}
	if result > 1000000000 {
		return -1
	}
	return result
}
