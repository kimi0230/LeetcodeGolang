package binarygap

// O(log n)
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

// https://wandbox.org/permlink/totZwDAbL1wCgsqt
func evil(x int) int {
	if x&(x+1) > 0 {
		return evil(x|(x>>1)) + 1
	} else {
		return 0
	}
}

func SolutionRecur(N int) int {
	for (N & 1) == 0 {
		N = N >> 1
	}
	return evil(N)
}
