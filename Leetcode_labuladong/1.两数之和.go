/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
// 用一個map, key紀錄number的值, value 紀錄number的index
// 遍歷整個nums,  判斷map中是否有 "target-number"存入map, 如果有就可以將 index從map取出, 並回傳
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

// @lc code=end

