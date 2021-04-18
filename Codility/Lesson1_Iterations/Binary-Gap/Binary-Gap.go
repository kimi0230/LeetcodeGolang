package binarygap

// O(N)
func Solution(N int) int {
	maxLen, curLen := 0, 0
	findOne := false
	for N > 0 {
		curBit := N & 1
		if curBit == 1 {

			curLen = 0
			findOne = true
		} else if curBit == 0 && findOne {
			curLen++
		}

		if curLen > maxLen {
			maxLen = curLen
		}
		N = N >> 1
	}
	return maxLen
}
