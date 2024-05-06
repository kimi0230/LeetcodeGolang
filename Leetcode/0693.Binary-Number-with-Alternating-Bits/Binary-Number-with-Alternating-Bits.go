package binarynumberwithalternatingbits

// 暴力解 O(n)
func hasAlternatingBits(n int) bool {
	for n > 0 {
		preBit := n & 1
		n = n / 2
		curBit := n & 1
		if curBit == preBit {
			return false
		}
	}
	return true
}

// 數學解
func hasAlternatingBits2(n int) bool {
	/*
		n=5
		n=				1 0 1
		n >> 1			0 1 0
		n^(n>>1)		1 1 1  (XOR 不同時得1)
		n               1 1 1
		n+1			  1 0 0 0
		n & (n+1)	    0 0 0

		n=7
		n=				1 1 1
		n >> 1			0 1 1
		n^(n>>1)		1 0 0  (XOR 不同時得1)
		n               1 0 0
		n+1			    1 0 1
		n & (n+1)	    1 0 0
	*/
	n = n ^ (n >> 1)
	result := n & (n + 1)
	return result == 0
}
