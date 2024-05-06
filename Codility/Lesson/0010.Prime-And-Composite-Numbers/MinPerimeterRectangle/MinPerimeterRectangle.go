package minperimeterrectangle

import (
	"math"
)

// O(sqrt(N))
func Solution(N int) int {
	if N <= 0 {
		return 0
	}

	min := math.MaxInt32
	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			perimeter := 2 * (i + N/i)
			min = int(math.Min(float64(min), float64(perimeter)))
		}

	}

	if min == math.MaxInt32 {
		return 0
	}
	return min
}

/*
O(N)
Task Score 60%
Correctness 100%
Performance 20%
*/
func Solution1(N int) int {
	if N <= 0 {
		return 0
	}

	min := math.MaxInt32
	for i := 1; i <= N; i++ {
		if N%i == 0 && i*i <= N {
			perimeter := 2 * (i + N/i)
			min = int(math.Min(float64(min), float64(perimeter)))
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}

// O(sqrt(N))
func Solution2(N int) int {
	if N <= 0 {
		return 0
	}
	pairs := make(map[int]int)
	i := 1
	for i*i <= N {
		if N%i == 0 {
			pairs[i] = N / i
		}
		i++
	}

	min := math.MaxInt32
	for i, v := range pairs {
		perimeter := 2 * (i + v)
		min = int(math.Min(float64(min), float64(perimeter)))

	}

	if min == math.MaxInt32 {
		return 0
	}
	return min
}
