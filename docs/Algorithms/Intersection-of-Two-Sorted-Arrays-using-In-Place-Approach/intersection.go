package intersection

func FindIntersection(A, B []int) []int {
	var i, j int = 0, 0
	result := []int{}

	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			i++
		} else if A[i] > B[j] {
			j++
		} else {
			result = append(result, A[i])
			i++
			j++
		}
	}
	return result
}
