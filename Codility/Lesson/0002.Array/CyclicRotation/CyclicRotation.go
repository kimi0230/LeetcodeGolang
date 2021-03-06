package cyclicrotation

func Solution(A []int, K int) []int {
	if K == 0 || len(A) <= 1 {
		return A
	}

	K = K % len(A)
	return append(A[len(A)-K:], A[:len(A)-K]...)
}

func Solution2(A []int, K int) []int {
	if K == 0 || len(A) <= 1 {
		return A
	}
	if K > len(A) {
		K = K % len(A)
	}
	return append(A[len(A)-K:], A[:len(A)-K]...)
}
