package findminimumdifferencebetweenanytwoelements

import (
	"math"
	"sort"
)

/*
https://yourbasic.org/golang/absolute-value-int-float/
http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
fmt.Println(abs(math.MinInt64)) // Output: -9223372036854775808

// 似乎比較快
func WithTwosComplement(n int64) int64 {
	y := n >> 63          // y ← x ⟫ 63
	return (n ^ y) - y    // (x ⨁ y) - y
}
*/
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func WithTwosComplement(n int64) int64 {
	y := n >> 63       // y ← x ⟫ 63
	return (n ^ y) - y // (x ⨁ y) - y
}

// O(n Log n)
func FindMinDiff(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	sort.Ints(nums) // O(n Log n)

	minSize := math.MaxInt32

	// O(n)
	for i := 1; i < len(nums); i++ {
		tmp := int(math.Abs(float64(nums[i] - nums[i-1])))
		if minSize > tmp {
			minSize = tmp
		}
	}

	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

func FindMinDiff2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	sort.Ints(nums)

	minSize := math.MaxInt32
	for i := 1; i < len(nums); i++ {
		tmp := abs(nums[i] - nums[i-1])
		if minSize > tmp {
			minSize = tmp
		}
	}

	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}
