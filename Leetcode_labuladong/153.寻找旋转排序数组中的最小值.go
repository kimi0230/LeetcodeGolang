// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := int(uint(left+right) >> 1)
		if nums[mid] < left {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

// @lc code=end

