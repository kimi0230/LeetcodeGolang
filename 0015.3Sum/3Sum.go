package threesum

import "sort"

// ThreeSumBurst : 暴力解 : O(n^3)
func ThreeSumBurst(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums) // O(n log n)
	for i := 0; i < len(nums); i++ {
		// 需要跟上一次不同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			// 需要跟上一次不同
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}

// 排序 + 雙指針法
func ThreeSumDoublePoint(nums []int) [][]int {
	result := [][]int{}
	return result
}
