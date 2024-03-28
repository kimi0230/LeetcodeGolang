// 時間複雜度: O(n)
// 空間複雜度: O(1)
/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		tmp := (right - left) * min(height[left], height[right])
		result = max(result, tmp)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

