# [0031. Find Minimum Difference Between Any Two Elements](https://www.geeksforgeeks.org/find-minimum-difference-pair/)

## 題目
Given an unsorted array, find the minimum difference between any pair in given array.
Examples :

Input  : {1, 5, 3, 19, 18, 25};
Output : 1
Minimum difference is between 18 and 19

Input  : {30, 5, 20, 9};
Output : 4
Minimum difference is between 5 and 9

Input  : {1, 19, -4, 31, 38, 25, 100};
Output : 5
Minimum difference is between 1 and -4

## 來源
* https://www.geeksforgeeks.org/find-minimum-difference-pair/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/GeeksforGeeks/SortingAlgorithms/0031.Find-Minimum-Difference-Between-Any-Two-Elements/Find-Minimum-Difference-Between-Any-Two-Elements.go

```go
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
```