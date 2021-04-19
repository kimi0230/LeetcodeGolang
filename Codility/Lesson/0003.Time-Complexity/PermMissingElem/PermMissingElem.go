package permmissingelem

func Solution(A []int) int {
	if len(A) < 1 {
		return 1
	}

	n := len(A) + 1
	predictSume := (n + 1) * n / 2

	var sum int
	for _, v := range A {
		sum += v
	}
	return predictSume - sum
}
