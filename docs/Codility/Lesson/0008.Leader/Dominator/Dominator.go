package Dominator

import (
	"math"
)

func Solution(A []int) int {
	mapInt := make(map[int]int, len(A))

	for _, v := range A {
		if _, ok := mapInt[v]; !ok {
			mapInt[v] = 1
		} else {
			mapInt[v]++
		}
	}

	maxCount := 0
	maxVal := 0
	for k, v := range mapInt {
		if v > maxCount {
			maxCount = v
			maxVal = k
		}
	}
	minIndex := -1
	for k, v := range A {
		if v == maxVal {
			minIndex = k
			break
		}
	}

	if maxCount > int(math.Floor(float64(len(A))/2.0)) {
		return minIndex
	} else {
		return -1
	}
}
