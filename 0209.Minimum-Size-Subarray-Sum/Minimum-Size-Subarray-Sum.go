package minimumsizesubarraysum

import "math"

// MinSubArrayLenBurst : 暴力解 O(n^2)
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
				minSize = min(minSize, j-i+1)
			}
		}
	}
	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

// 滑動視窗 O(n)
func MinSubArrayLenSlidingWindow(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}
	minSize := math.MaxInt32

	start, end, sum := 0, 0, 0
	for end < lens {
		sum += nums[end]
		for sum >= target {
			minSize = min(minSize, end-start+1)
			sum -= nums[start]
			start++
		}
		end++

	}

	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
