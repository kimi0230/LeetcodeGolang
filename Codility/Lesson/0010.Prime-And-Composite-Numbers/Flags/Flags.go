package Flags

import (
	"math"
)

func Solution(A []int) int {
	var peaksFlag []int

	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaksFlag = append(peaksFlag, i)
		}
	}

	if len(peaksFlag) == 0 {
		return 0
	}
	if len(peaksFlag) == 1 {
		return 1
	}

	maxFlag := int(math.Pow(float64(peaksFlag[len(peaksFlag)-1]-peaksFlag[0]), 0.5) + 1)

	for i := maxFlag; i > 1; i-- {
		addressFlag := []int{peaksFlag[0]}
		for _, val := range peaksFlag[1:] {
			if val-addressFlag[len(addressFlag)-1] >= i {
				addressFlag = append(addressFlag, val)
				if len(addressFlag) >= i {
					return i
				}
			}
		}
	}

	return 1
}
