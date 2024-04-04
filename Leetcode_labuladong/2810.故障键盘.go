// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=2810 lang=golang
 *
 * [2810] 故障键盘
 */

// @lc code=start
func finalString(s string) string {
	queue := [2][]rune{} // 0:向左, 1:向右
	dir := 1
	for _, c := range s {
		if c == 'i' {
			dir ^= 1
		} else {
			queue[dir] = append(queue[dir], c)
		}
	}
	slices.Reverse(queue[dir^1])
	return string(append(queue[dir^1], queue[dir]...))
}

// @lc code=end

