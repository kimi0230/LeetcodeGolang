package minimumsizesubarraysum

import "math"

func MinSubArrayLenBurst(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}

	minSize := math.MaxInt32
	for i := 0; i < lens; i++ {
		sum := 0
		for j := i; j < lens; j++ {
			sum += nums[j]
			if sum >= target {
				if minSize > (j - i + 1) {
					minSize = j - i + 1
					break
				}
			}
		}
	}
	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}
