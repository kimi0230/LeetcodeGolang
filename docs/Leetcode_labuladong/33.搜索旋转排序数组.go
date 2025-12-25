// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right) >> 1)

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			// mid在上半區

			if nums[left] <= target && target <= nums[mid] {
				// target 在left跟min中間
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// mid在下半區
			if nums[mid] <= target && target <= nums[right] {
				// target 在min跟right中間
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return left
	} else if nums[right] == target {
		return right
	} else {
		return -1
	}
}

// @lc code=end

