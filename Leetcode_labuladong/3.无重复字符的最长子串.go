// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (39.64%)
 * Likes:    10085
 * Dislikes: 0
 * Total Accepted:    2.8M
 * Total Submissions: 7.1M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	// m := make(map[byte]bool, len(s))
	// 使用slice優化
	m := [256]bool{}
	left, right := 0, 0
	maxLen := 0
	for left < slength {
		if ok := m[s[right]]; ok {
			m[s[left]] = false
			left += 1
		} else {
			m[s[right]] = true
			right += 1
		}
		if maxLen < (right - left) {
			maxLen = right - left
		}
		if (left+maxLen) >= slength || right >= slength {
			break
		}
	}
	return maxLen
}

// @lc code=end

