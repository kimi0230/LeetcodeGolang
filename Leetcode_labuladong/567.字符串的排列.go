// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode.cn/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (44.95%)
 * Likes:    998
 * Dislikes: 0
 * Total Accepted:    287.5K
 * Total Submissions: 639.1K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
 *
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "ab" s2 = "eidbaooo"
 * 输出：true
 * 解释：s2 包含 s1 的排列之一 ("ba").
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1= "ab" s2 = "eidboaoo"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length, s2.length <= 10^4
 * s1 和 s2 仅包含小写字母
 *
 *
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	need, windows := make(map[byte]int), make(map[byte]int)
	match := 0
	for _, v := range s1 {
		need[byte(v)]++
	}

	for right < len(s2) {
		cRight := s2[right]
		windows[cRight]++
		right++
		if windows[cRight] == need[cRight] {
			match++
		}

		if match == len(need) {
			return true
		}

		if right-left >= len(s1) {
			cLeft := s2[left]

			// 如果左指針字符出現次數達到目標，減少 match
			if windows[cLeft] == need[cLeft] {
				match--
			}
			// 移動左指針
			windows[cLeft]--
			left++
		}
	}
	return false
}

// @lc code=end

