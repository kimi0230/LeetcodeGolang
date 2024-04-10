package longestSubstringwithoutrepeatingcharacters

// LengthOfLongestSubstring 暴力解
func LengthOfLongestSubstring(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	tmpLen := 1
	var maxLen = 1

	for i := 1; i < slength; i++ {
		// 往前找前幾個視窗
		j := i - tmpLen

		for ; j < i; j++ {
			if s[j] == s[i] { // 如果相同，那麼和S[J]到S[I-1]中間的肯定不相同，所以可以直接計算得到
				tmpLen = i - j
				break
			}
		}

		if j == i { // 都不相同
			tmpLen++
		}

		if tmpLen > maxLen {
			maxLen = tmpLen
		}
	}

	return maxLen
}

// LengthOfLongestSubstringMap 用map 紀錄是否重複.
// O(n)
func LengthOfLongestSubstringMap(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	charMap := make(map[byte]bool)
	maxLen, left, right := 0, 0, 0

	for left < slength {
		if ok := charMap[s[right]]; ok {
			// 有找到
			charMap[s[left]] = false
			left++
		} else {
			charMap[s[right]] = true
			right++
		}
		if maxLen < right-left {
			maxLen = right - left
		}
		if (left+maxLen) >= slength || right >= len(s) {
			break
		}
	}

	return maxLen
}

// LengthOfLongestSubstringBit 用map效能不好時可使用數組改善
func LengthOfLongestSubstringBit(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	// ASCII 0~255
	charMap := [256]bool{}
	maxLen, left, right := 0, 0, 0
	for left < slength {
		if ok := charMap[s[right]]; ok {
			// 有找到
			charMap[s[left]] = false
			left++
		} else {
			charMap[s[right]] = true
			right++
		}

		if maxLen < right-left {
			maxLen = right - left
		}
		if left+maxLen >= slength || right >= len(s) {
			break
		}
	}
	return maxLen
}
