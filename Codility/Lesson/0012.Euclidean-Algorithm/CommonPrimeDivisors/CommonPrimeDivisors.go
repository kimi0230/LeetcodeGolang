package commonprimedivisors

/*
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
*/
func gcd(N int, M int) int {
	if N%M == 0 {
		return M
	} else {
		return gcd(M, N%M)
	}
}

func Solution(A []int, B []int) int {
	result := 0
	for i := 0; i < len(A); i++ {
		if A[i] == B[i] {
			result++
			continue
		}
		// 先判斷兩數的最大公因數,
		abGcd := gcd(A[i], B[i])

		// 再判斷兩個數是含有最大公因數沒有的因子
		a := A[i] / abGcd
		aGcd := gcd(a, abGcd)
		for aGcd != 1 {
			// 還有其他因子
			a = a / aGcd
			aGcd = gcd(aGcd, a)
		}

		b := B[i] / abGcd
		bGcd := gcd(b, abGcd)
		for bGcd != 1 {
			b = b / bGcd
			bGcd = gcd(bGcd, b)
		}
		if a == b {
			result++
		}
	}
	return result
}
