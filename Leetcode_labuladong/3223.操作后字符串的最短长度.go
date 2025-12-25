/*
 * @lc app=leetcode.cn id=3223 lang=golang
 * @lcpr version=30104
 *
 * [3223] 操作后字符串的最短长度
 */

// @lc code=start
func minimumLength(s string) int {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	length := 0
	for _, x := range cnt {
		if x > 0 {
			if x%2 == 0 {
				length += 2
			} else {
				length += 1
			}
		}
	}
	return length
}

// @lc code=end

/*
// @lcpr case=start
// "abaacbcbb"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n
// @lcpr case=end

*/
