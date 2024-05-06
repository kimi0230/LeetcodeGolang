// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	result := []int{}

	// first index
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right) >> 1) // [0,1] => 0
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid
		}
	}
	if left == right && nums[left] == target {
		result = append(result, left)
	} else {
		result = append(result, -1)
	}

	// last index
	left, right = 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right+1) >> 1) // [0,1] => 1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if left == right && nums[left] == target {
		result = append(result, left)
	} else {
		result = append(result, -1)
	}
	return result
}

// @lc code=end

